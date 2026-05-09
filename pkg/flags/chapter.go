package flags

import (
	"github.com/mniak/biblia/pkg/bible"
)

// ChapterFlag implements pflag.Value for parsing chapter references
// Supports formats like: "Genesis1", "GEN1", "Gn1", "MAT5", "IJoao3"
type ChapterFlag struct {
	value   bible.ChapterRef
	aliases map[string]bible.BookCode
	isSet   bool
}

// NewChapterFlag creates a new ChapterFlag with the given aliases
func NewChapterFlag(aliases map[string]bible.BookCode) *ChapterFlag {
	return &ChapterFlag{
		aliases: aliases,
	}
}

// String implements pflag.Value
func (f *ChapterFlag) String() string {
	if !f.isSet {
		return ""
	}
	return f.value.String()
}

// Set implements pflag.Value
func (f *ChapterFlag) Set(value string) error {
	ref, err := bible.ParseChapterRef(value, f.aliases)
	if err != nil {
		return err
	}
	f.value = ref
	f.isSet = true
	return nil
}

// Type implements pflag.Value
func (f *ChapterFlag) Type() string {
	return "chapter"
}

// Value returns the parsed ChapterRef
func (f *ChapterFlag) Value() bible.ChapterRef {
	return f.value
}

// IsSet returns true if a value has been set
func (f *ChapterFlag) IsSet() bool {
	return f.isSet
}
