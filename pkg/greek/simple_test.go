package greek

import (
	"fmt"
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
)

var _ bible.Transliterator = SimpleTransliterator()

func TestSimple_Examples(t *testing.T) {
	testData := map[string]string{
		"Ὧδε":       "Hode",
		"ἡ":         "he",
		"ὑπομονὴ":   "hypomone",
		"τῶν":       "ton",
		"ἁγίων":     "hagion",
		"ἐστίν·":    "estin",
		"οἱ":        "hoi",
		"τηροῦντες": "terountes",
		"τὰς":       "tas",
		"ἐντολὰς":   "entolas",
		"τοῦ":       "tou",
		"θεοῦ":      "theou",
		"καὶ":       "kai",
		"τὴν":       "ten",
		"πίστιν":    "pistin",
		"Ἰησοῦ":     "Iesou",
	}

	trans := SimpleTransliterator()

	for word, expected := range testData {
		t.Run(fmt.Sprintf("'%s'→'%s'", word, expected), func(t *testing.T) {
			actual := trans.TransliterateWord(word)
			assert.Equal(t, expected, actual)
		})
	}
}
