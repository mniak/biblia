package transliterators

import (
	"fmt"
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
)

var _ bible.Transliterator = PhoneticPTBR()

func TestPhoneticPTBR_Simple(t *testing.T) {
	testData := map[string]string{
		"בְּרֵאשִׁ֖ית": "bereshit",
		"בָּרָ֣א":      "bara",
		"אֱלֹהִ֑ים":    "elohim",
		"אֵ֥ת":         "et",
		"הַשָּׁמַ֖יִם": "hashmaim",
		"וְאֵ֥ת":       "we-et",
		"הָאָֽרֶץ׃":    "haarets",
		"תֹ֙הוּ֙":      "tohu",
		"א֑וֹר":        "or",
	}

	trans := PhoneticPTBR()

	for word, expected := range testData {
		t.Run(fmt.Sprintf("'%s'→'%s'", word, expected), func(t *testing.T) {
			actual := trans.TransliterateWord(word)

			assert.Equal(t, expected, actual)
		})
	}
}
