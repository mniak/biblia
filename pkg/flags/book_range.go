package flags

import (
	"fmt"
	"strings"

	"github.com/mniak/biblia/pkg/bible"
)

// BookIndexParser returns a parser function for BookIndex that accepts:
//   - Book codes: "GEN", "EXO", "MAT"
//   - Aliases from the provided map: "Genesis", "Gn", etc.
func BookIndexParser(aliases map[string]bible.BookCode) func(string) (bible.BookIndex, error) {
	return func(s string) (bible.BookIndex, error) {
		s = strings.TrimSpace(s)

		// Try parsing as book code
		code := bible.BookCode(strings.ToUpper(s))
		if idx := code.Index(); idx != 0 {
			return idx, nil
		}

		// Try aliases (case-insensitive)
		if aliases != nil {
			lower := strings.ToLower(s)
			for alias, bookCode := range aliases {
				if strings.ToLower(alias) == lower {
					return bookCode.Index(), nil
				}
			}
		}

		return 0, fmt.Errorf("unknown book: %s", s)
	}
}

// NewBookIndexRangeFlag creates a RangeFlag[bible.BookIndex] with book code parsing
func NewBookIndexRangeFlag(aliases map[string]bible.BookCode) *RangeFlag[bible.BookIndex] {
	return NewRangeFlag(BookIndexParser(aliases))
}
