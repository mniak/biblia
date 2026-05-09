package flags

import (
	"github.com/mniak/biblia/pkg/bible"
)

// BookFlag implements pflag.Value for parsing biblical books
type BookFlag struct {
	value   bible.BookCode
	aliases map[string]bible.BookCode
	isSet   bool
}

// NewBookFlag creates a new BookFlag with the given aliases
func NewBookFlag(aliases map[string]bible.BookCode) *BookFlag {
	return &BookFlag{
		aliases: aliases,
	}
}

// String implements pflag.Value
func (f *BookFlag) String() string {
	if !f.isSet {
		return ""
	}
	return f.value.String()
}

// Set implements pflag.Value
func (f *BookFlag) Set(value string) error {
	code, err := bible.ParseBookCode(value, f.aliases)
	if err != nil {
		return err
	}
	f.value = code
	f.isSet = true
	return nil
}

// Type implements pflag.Value
func (f *BookFlag) Type() string {
	return "book"
}

// Value returns the parsed BookCode
func (f *BookFlag) Value() bible.BookCode {
	return f.value
}

// IsSet returns true if a value has been set
func (f *BookFlag) IsSet() bool {
	return f.isSet
}
