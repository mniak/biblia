package flagutils

import (
	"fmt"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/romanizers/academic"
)

func Romanizers(lang bible.Language) map[string]bible.Romanizer {
	result := make(map[string]bible.Romanizer)
	if lang == bible.LangHebrew || lang == bible.LangAny {
		result["academic-hebrew"] = academic.NewHebrewRomanizer()
	}
	if lang == bible.LangGreek || lang == bible.LangAny {
	}
	return result
}

func Romanizer(lang bible.Language, romanizer string) (bible.Romanizer, error) {
	result, found := Romanizers(lang)[romanizer]
	if !found {
		return nil, fmt.Errorf("invalid romanizer: %s", result)
	}

	return result, nil
}
