package bible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllBookInfos(t *testing.T) {
	books := AllOldTestamentBooks()
	for _, book := range books {
		chapters := book.ChapterCount()
		assert.Greater(t, chapters, 0)

		for chapter := 1; chapter <= chapters; chapter++ {
			verses := book.VerseCount(chapter)
			assert.Greater(t, verses, 0)
		}
	}
}
