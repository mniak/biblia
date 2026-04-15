package flags

import (
	"github.com/mniak/biblia/pkg/bible"
)

// NewBookIndexRangeFlag creates a RangeFlag[bible.BookIndex] with book code parsing
func NewBookIndexRangeFlag(aliases map[string]bible.BookCode) *RangeFlag[bible.BookIndex] {
	return NewRangeFlag(func(s string) (bible.BookIndex, error) {
		return bible.ParseBookIndex(s, aliases)
	})
}
