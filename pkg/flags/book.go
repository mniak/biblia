package flags

import (
	"fmt"
	"strings"

	"github.com/mniak/biblia/pkg/bible"
)

// BookInfo contains information about a biblical book
type BookInfo struct {
	Code    bible.BookCode
	Aliases []string
}

// BookFlag implements pflag.Value for parsing biblical books
type BookFlag struct {
	books    []BookInfo
	selected *BookInfo
}

// NewBookFlag creates a new BookFlag with the given book definitions
func NewBookFlag(books []BookInfo) *BookFlag {
	return &BookFlag{
		books: books,
	}
}

// String implements pflag.Value
func (f *BookFlag) String() string {
	if f.selected == nil {
		return ""
	}
	return f.selected.Code.String()
}

// Set implements pflag.Value
func (f *BookFlag) Set(value string) error {
	value = strings.ToLower(strings.TrimSpace(value))

	for i := range f.books {
		book := &f.books[i]

		// Check by code
		if strings.ToLower(book.Code.String()) == value {
			f.selected = book
			return nil
		}

		// Check by aliases
		for _, alias := range book.Aliases {
			if strings.ToLower(alias) == value {
				f.selected = book
				return nil
			}
		}
	}

	return fmt.Errorf("unknown book: %s", value)
}

// Type implements pflag.Value
func (f *BookFlag) Type() string {
	return "book"
}

// Value returns the selected book, or nil if none selected
func (f *BookFlag) Value() *BookInfo {
	return f.selected
}

// IsSet returns true if a book has been selected
func (f *BookFlag) IsSet() bool {
	return f.selected != nil
}
