package transliterators

import (
	"fmt"
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
)

var _ bible.Transliterator = Academic()

func TestAcademic_Simple(t *testing.T) {
	testData := map[string]string{
		"בְּרֵאשִׁ֖ית": "b'rēʾšîṯ",
		"בָּרָ֣א":      "bārāʾ",
		"אֱלֹהִ֑ים":    "ʾĕlohîm",
		"אֵ֥ת":         "ʾēṯ",
		"הַשָּׁמַ֖יִם": "haššāmayim",
		"וְאֵ֥ת":       "w'ʾēṯ",
		"הָאָֽרֶץ׃":    "hāʾāreṣ",
		"תֹ֙הוּ֙":      "ṯohû",
		"א֑וֹר":        "ʾôr",
	}

	transliterator := Academic()

	for word, expected := range testData {
		t.Run(fmt.Sprintf("'%s'→'%s'", word, expected), func(t *testing.T) {
			actual := transliterator.TransliterateWord(word)

			assert.Equal(t, expected, actual)
		})
	}
}
