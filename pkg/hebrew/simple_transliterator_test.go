package hebrew

import (
	"fmt"
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
)

var _ bible.Transliterator = SimpleTransliterator()

func TestPortuguese_Simple(t *testing.T) {
	testData := map[string]string{
		"בְּרֵאשִׁ֖ית": "bᵉreshit",
		"בָּרָ֣א":      "bara",
		"אֱלֹהִ֑ים":    "elohim",
		"אֵ֥ת":         "et",
		"הַשָּׁמַ֖יִם": "hashamayim",
		"וְאֵ֥ת":       "wᵉet",
		"הָאָֽרֶץ׃":    "haarets",
		"תֹ֙הוּ֙":      "tohu",
		"א֑וֹר":        "or",
	}
	transliterator := SimpleTransliterator()

	for word, expected := range testData {
		t.Run(fmt.Sprintf("'%s'→'%s'", word, expected), func(t *testing.T) {
			actual := transliterator.TransliterateWord(word)
			assert.Equal(t, expected, actual)
		})
	}
}
