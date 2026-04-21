package bible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChapterRange_Chapters(t *testing.T) {
	r := NewChapterRangeFromTo(Leviticus, 10, 12)

	var chapters []ChapterRef
	for ch := range r.Chapters() {
		chapters = append(chapters, ch)
	}

	assert.Equal(t, []ChapterRef{
		{Leviticus, 10},
		{Leviticus, 11},
		{Leviticus, 12},
	}, chapters)
}

func TestChapterRange_SingleChapter(t *testing.T) {
	r := NewChapterRange(Psalms, 23)

	chapters := r.Collect()

	assert.Equal(t, []ChapterRef{{Psalms, 23}}, chapters)
}

func TestChapterRange_Contains(t *testing.T) {
	r := NewChapterRangeFromTo(Genesis, 1, 11)

	assert.True(t, r.Contains(ChapterRef{Genesis, 1}))
	assert.True(t, r.Contains(ChapterRef{Genesis, 5}))
	assert.True(t, r.Contains(ChapterRef{Genesis, 11}))
	assert.False(t, r.Contains(ChapterRef{Genesis, 12}))
	assert.False(t, r.Contains(ChapterRef{Exodus, 1}))
}

func TestChapterRange_ContainsChapter(t *testing.T) {
	r := NewChapterRangeFromTo(Genesis, 1, 11)

	assert.True(t, r.ContainsChapter(1))
	assert.True(t, r.ContainsChapter(11))
	assert.False(t, r.ContainsChapter(0))
	assert.False(t, r.ContainsChapter(12))
}

func TestChapterRange_String(t *testing.T) {
	t.Run("single chapter", func(t *testing.T) {
		r := NewChapterRange(Psalms, 23)
		assert.Equal(t, "PSA23", r.String())
	})

	t.Run("range", func(t *testing.T) {
		r := NewChapterRangeFromTo(Leviticus, 10, 12)
		assert.Equal(t, "LEV10-12", r.String())
	})
}

func TestChapterRange_Count(t *testing.T) {
	t.Run("single chapter", func(t *testing.T) {
		r := NewChapterRange(Psalms, 23)
		assert.Equal(t, 1, r.Count())
	})

	t.Run("range", func(t *testing.T) {
		r := NewChapterRangeFromTo(Genesis, 1, 11)
		assert.Equal(t, 11, r.Count())
	})
}
