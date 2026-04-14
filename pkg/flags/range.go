package flags

import (
	"cmp"
	"fmt"
	"iter"
	"strconv"
	"strings"
)

// Rangeable represents types that can be used in range expressions
type Rangeable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// segment represents a single value or range of values
type segment[T Rangeable] struct {
	start T
	end   T
}

// RangeFlag implements pflag.Value for parsing ranges of values
// Supports:
//   - Single values: "5"
//   - Ranges: "2-7" (inclusive)
//   - Multiple values: "2,5,7"
//   - Combined: "1-3,11-14,20"
type RangeFlag[T Rangeable] struct {
	segments []segment[T]
	parser   func(string) (T, error)
}

// NewRangeFlag creates a new RangeFlag with a custom parser
func NewRangeFlag[T Rangeable](parser func(string) (T, error)) *RangeFlag[T] {
	return &RangeFlag[T]{
		parser: parser,
	}
}

// NewIntRangeFlag creates a RangeFlag for int values
func NewIntRangeFlag() *RangeFlag[int] {
	return NewRangeFlag(func(s string) (int, error) {
		return strconv.Atoi(s)
	})
}

// String implements pflag.Value
func (f *RangeFlag[T]) String() string {
	if len(f.segments) == 0 {
		return ""
	}

	var parts []string
	for _, seg := range f.segments {
		if seg.start == seg.end {
			parts = append(parts, fmt.Sprintf("%v", seg.start))
		} else {
			parts = append(parts, fmt.Sprintf("%v-%v", seg.start, seg.end))
		}
	}
	return strings.Join(parts, ",")
}

// Set implements pflag.Value
func (f *RangeFlag[T]) Set(value string) error {
	f.segments = nil

	parts := strings.Split(value, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		seg, err := f.parseSegment(part)
		if err != nil {
			return err
		}
		f.segments = append(f.segments, seg)
	}

	return nil
}

func (f *RangeFlag[T]) parseSegment(s string) (segment[T], error) {
	if idx := strings.Index(s, "-"); idx > 0 {
		startStr := strings.TrimSpace(s[:idx])
		endStr := strings.TrimSpace(s[idx+1:])

		start, err := f.parser(startStr)
		if err != nil {
			return segment[T]{}, fmt.Errorf("invalid range start %q: %w", startStr, err)
		}

		end, err := f.parser(endStr)
		if err != nil {
			return segment[T]{}, fmt.Errorf("invalid range end %q: %w", endStr, err)
		}

		if cmp.Compare(start, end) > 0 {
			return segment[T]{}, fmt.Errorf("invalid range: start %v is greater than end %v", start, end)
		}

		return segment[T]{start: start, end: end}, nil
	}

	val, err := f.parser(s)
	if err != nil {
		return segment[T]{}, fmt.Errorf("invalid value %q: %w", s, err)
	}

	return segment[T]{start: val, end: val}, nil
}

// Type implements pflag.Value
func (f *RangeFlag[T]) Type() string {
	return "range"
}

// Values returns an iterator over all values in the range
func (f *RangeFlag[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seg := range f.segments {
			for v := seg.start; v <= seg.end; v++ {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Collect returns all values as a slice
func (f *RangeFlag[T]) Collect() []T {
	var result []T
	for v := range f.Values() {
		result = append(result, v)
	}
	return result
}

// Contains checks if a value is within any of the ranges
func (f *RangeFlag[T]) Contains(value T) bool {
	for _, seg := range f.segments {
		if cmp.Compare(value, seg.start) >= 0 && cmp.Compare(value, seg.end) <= 0 {
			return true
		}
	}
	return false
}

// IsSet returns true if any values have been set
func (f *RangeFlag[T]) IsSet() bool {
	return len(f.segments) > 0
}
