package flagutils

import (
	"fmt"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/wlc"
)

func TestamentLoaders(lang bible.Language) map[string]bible.TestamentLoader {
	result := make(map[string]bible.TestamentLoader)
	if lang == bible.LangHebrew || lang == bible.LangAny {
		result["wlc"] = wlc.Loader("sources/UnicodeXML_Westminster_Leningrad_Codex/Tanach")
	}
	if lang == bible.LangGreek || lang == bible.LangAny {
	}
	return result
}

func TestamentLoader(lang bible.Language, source string) (bible.TestamentLoader, error) {
	result, found := TestamentLoaders(lang)[source]
	if !found {
		return nil, fmt.Errorf("invalid loader: %s", result)
	}

	return result, nil
}
