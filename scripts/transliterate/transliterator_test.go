package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var academicTransliterator Transliterator = &_Transliterator{}

func TestAcademic_Simple(t *testing.T) {
	testData := map[string]string{
		"בְּרֵאשִׁ֖ית": "b'reʾšîṯ",
		// "בָּרָ֣א":      "",
		// "אֱלֹהִ֑ים":    "",
		// "אֵ֥ת":         "",
		// "הַשָּׁמַ֖יִם": "",
		// "וְאֵ֥ת":       "",
		// "הָאָֽרֶץ׃":    "",
	}

	for word, expected := range testData {
		t.Run(fmt.Sprintf("'%s' -> '%s'", word, expected), func(t *testing.T) {
			actual := academicTransliterator.Transliterate(word)

			assert.Equal(t, expected, actual)
		})
	}
}
