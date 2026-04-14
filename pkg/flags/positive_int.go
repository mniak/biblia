package flags

import (
	"fmt"
	"strconv"
)

// PositiveIntFlag implements pflag.Value for parsing positive integers
type PositiveIntFlag struct {
	value int
	isSet bool
}

// NewPositiveIntFlag creates a new PositiveIntFlag with a default value
func NewPositiveIntFlag(defaultValue int) *PositiveIntFlag {
	return &PositiveIntFlag{
		value: defaultValue,
		isSet: false,
	}
}

// String implements pflag.Value
func (f *PositiveIntFlag) String() string {
	return strconv.Itoa(f.value)
}

// Set implements pflag.Value
func (f *PositiveIntFlag) Set(value string) error {
	n, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("invalid integer: %s", value)
	}

	if n <= 0 {
		return fmt.Errorf("value must be positive: %d", n)
	}

	f.value = n
	f.isSet = true
	return nil
}

// Type implements pflag.Value
func (f *PositiveIntFlag) Type() string {
	return "positive-int"
}

// Value returns the integer value
func (f *PositiveIntFlag) Value() int {
	return f.value
}

// IsSet returns true if a value has been explicitly set
func (f *PositiveIntFlag) IsSet() bool {
	return f.isSet
}
