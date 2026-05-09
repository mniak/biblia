package flags

import (
	"github.com/mniak/biblia/pkg/bible"
)

// VerseFlag implements pflag.Value for parsing verse references
// Supports formats like: "Genesis1:1", "GEN1:1", "Gn1:1", "MAT5:3", "iJoao4:8"
type VerseFlag struct {
	value   bible.VerseRef
	aliases map[string]bible.BookCode
	isSet   bool
}

// NewVerseFlag creates a new VerseFlag with the given aliases
func NewVerseFlag(aliases map[string]bible.BookCode) *VerseFlag {
	return &VerseFlag{
		aliases: aliases,
	}
}

// String implements pflag.Value
func (f *VerseFlag) String() string {
	if !f.isSet {
		return ""
	}
	return f.value.String()
}

// Set implements pflag.Value
func (f *VerseFlag) Set(value string) error {
	ref, err := bible.ParseVerseRef(value, f.aliases)
	if err != nil {
		return err
	}
	f.value = ref
	f.isSet = true
	return nil
}

// Type implements pflag.Value
func (f *VerseFlag) Type() string {
	return "verse"
}

// Value returns the parsed VerseRef
func (f *VerseFlag) Value() bible.VerseRef {
	return f.value
}

// IsSet returns true if a value has been set
func (f *VerseFlag) IsSet() bool {
	return f.isSet
}
