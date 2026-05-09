package bible

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerseRange_Verses(t *testing.T) {
	r := NewVerseRangeFromTo(Exodus, 20, 8, 11)

	var verses []VerseRef
	for v := range r.Verses() {
		verses = append(verses, v)
	}

	assert.Equal(t, []VerseRef{
		{Exodus, 20, 8},
		{Exodus, 20, 9},
		{Exodus, 20, 10},
		{Exodus, 20, 11},
	}, verses)
}

func TestVerseRange_SingleVerse(t *testing.T) {
	r := NewVerseRange(John, 3, 16)

	verses := r.Collect()

	assert.Equal(t, []VerseRef{{John, 3, 16}}, verses)
}

func TestVerseRange_Contains(t *testing.T) {
	r := NewVerseRangeFromTo(Matthew, 5, 3, 12)

	assert.True(t, r.Contains(VerseRef{Matthew, 5, 3}))
	assert.True(t, r.Contains(VerseRef{Matthew, 5, 7}))
	assert.True(t, r.Contains(VerseRef{Matthew, 5, 12}))
	assert.False(t, r.Contains(VerseRef{Matthew, 5, 2}))
	assert.False(t, r.Contains(VerseRef{Matthew, 5, 13}))
	assert.False(t, r.Contains(VerseRef{Matthew, 6, 5}))
	assert.False(t, r.Contains(VerseRef{Mark, 5, 5}))
}

func TestVerseRange_ContainsVerse(t *testing.T) {
	r := NewVerseRangeFromTo(Matthew, 5, 3, 12)

	assert.True(t, r.ContainsVerse(3))
	assert.True(t, r.ContainsVerse(12))
	assert.False(t, r.ContainsVerse(2))
	assert.False(t, r.ContainsVerse(13))
}

func TestVerseRange_String(t *testing.T) {
	t.Run("single verse", func(t *testing.T) {
		r := NewVerseRange(John, 3, 16)
		assert.Equal(t, "JHN3:16", r.String())
	})

	t.Run("range", func(t *testing.T) {
		r := NewVerseRangeFromTo(Exodus, 20, 8, 11)
		assert.Equal(t, "EXO20:8-11", r.String())
	})
}

func TestVerseRange_Count(t *testing.T) {
	t.Run("single verse", func(t *testing.T) {
		r := NewVerseRange(John, 3, 16)
		assert.Equal(t, 1, r.Count())
	})

	t.Run("range", func(t *testing.T) {
		r := NewVerseRangeFromTo(Exodus, 20, 8, 11)
		assert.Equal(t, 4, r.Count())
	})
}

func TestVerseRange_ChapterRef(t *testing.T) {
	r := NewVerseRangeFromTo(Exodus, 20, 8, 11)

	assert.Equal(t, ChapterRef{Exodus, 20}, r.ChapterRef())
}

func TestPredefinedVerseRanges(t *testing.T) {
	tests := []struct {
		name    string
		r       VerseRange
		count   int
		book    BookCode
		chapter int
	}{
		{name: "TenCommandments", r: TenCommandments, count: 17, book: Exodus, chapter: 20},
		{name: "LordsPrayer", r: LordsPrayer, count: 5, book: Matthew, chapter: 6},
		{name: "Beatitudes", r: Beatitudes, count: 10, book: Matthew, chapter: 5},
		{name: "LoveChapter", r: LoveChapter, count: 13, book: ICorinthians, chapter: 13},
		{name: "John316", r: John316, count: 1, book: John, chapter: 3},
		{name: "Psalm23", r: Psalm23, count: 6, book: Psalms, chapter: 23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.count, tt.r.Count())
			assert.Equal(t, tt.book, tt.r.Book)
			assert.Equal(t, tt.chapter, tt.r.Chapter)
		})
	}
}
