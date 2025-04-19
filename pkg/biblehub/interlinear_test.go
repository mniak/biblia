package biblehub_test

import (
	"testing"

	"github.com/mniak/biblia/internal/test"
	"github.com/mniak/biblia/pkg/biblehub"
	"github.com/mniak/biblia/pkg/biblehub/biblehubtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInterlinearChapter_Daniel2(t *testing.T) {
	ex := biblehub.Scraper{
		Downloader: biblehubtest.FakeDownloader,
	}

	ch, err := ex.GetInterlinearChapter(biblehub.ChapterID{
		Book:    "daniel",
		Chapter: 2,
	})
	require.NoError(t, err)
	assert.Equal(t, biblehub.Hebrew, ch.Language)

	assert.Equal(t, "Daniel 2", ch.Title)
	assert.Equal(t, "daniel", ch.Book)
	assert.Equal(t, 2, ch.Chapter)
	require.NotNil(t, ch.Next)
	assert.Equal(t, "daniel", ch.Next.Book)
	assert.Equal(t, 3, ch.Next.Chapter)
	require.Len(t, ch.Verses, 49)

	v49 := ch.Verses[48]
	require.Len(t, v49.Words, 30)
	assert.Equal(t, 49, v49.Number)

	firstWord := v49.Words[0]
	assert.Equal(t, "וְדָנִיֵּאל֙", firstWord.Original)
	assert.Equal(t, "wə·ḏā·nî·yêl", firstWord.Transliteration)
	assert.Equal(t, "1841", firstWord.StrongsNumber)
	test.AssertEqualTrim(t,
		`Strong's Hebrew 1841: Daniel = God is my judge
		1) the 4th of the greater prophets, taken as hostage in the first deportation to Babylon, because of the gift of God of the interpretation of dreams, he became the 2nd in command of the Babylon empire and lasted through the end of the Babylonian empire and into the Persian empire. His prophecies are the key to the understanding of end time events. Noted for his purity and holiness by contemporary prophet, Ezekiel 
		1a) also, 'Belteshazzar' ( H01095 or H01096)`,
		firstWord.StrongsText)
	assert.Equal(t, "Conj‑w | N‑proper‑ms", firstWord.SyntaxCode)
	assert.Equal(t, "Conjunctive waw :: Noun - proper - masculine singular", firstWord.SyntaxDescription)
	assert.Equal(t, "And Daniel", firstWord.English)

	previousToLastWord := v49.Words[len(v49.Words)-2]
	assert.Equal(t, "מַלְכָּֽא׃", previousToLastWord.Original)
	assert.Equal(t, "mal·kā.", previousToLastWord.Transliteration)
	assert.Equal(t, "4430", previousToLastWord.StrongsNumber)
	test.AssertEqualTrim(t, `Strong's Hebrew 4430: 1) king`, previousToLastWord.StrongsText)
	assert.Equal(t, "N‑msd", previousToLastWord.SyntaxCode)
	assert.Equal(t, "Noun - masculine singular determinate", previousToLastWord.SyntaxDescription)
	assert.Equal(t, "of king the", previousToLastWord.English)

	lastWord := v49.Words[len(v49.Words)-1]
	assert.Equal(t, "פ", lastWord.Original)
	assert.Equal(t, "p̄", lastWord.Transliteration)
	assert.Equal(t, "Punc", lastWord.SyntaxCode)
	assert.Equal(t, "Punctuation", lastWord.SyntaxDescription)
	assert.Equal(t, "-", lastWord.English)
	assert.Empty(t, lastWord.StrongsNumber)
	assert.Empty(t, lastWord.StrongsText)
}

func TestGetInterlinearChapter_Revelation13(t *testing.T) {
	ex := biblehub.Scraper{
		Downloader: biblehubtest.FakeDownloader,
	}

	ch, err := ex.GetInterlinearChapter(biblehub.ChapterID{
		Book:    "revelation",
		Chapter: 13,
	})
	require.NoError(t, err)
	assert.Equal(t, biblehub.Greek, ch.Language)

	assert.Equal(t, "Revelation 13", ch.Title)
	assert.Equal(t, "revelation", ch.Book)
	assert.Equal(t, 13, ch.Chapter)
	require.NotNil(t, ch.Next)
	assert.Equal(t, "revelation", ch.Next.Book)
	assert.Equal(t, 14, ch.Next.Chapter)
	require.Len(t, ch.Verses, 18)

	v18 := ch.Verses[17]
	require.Len(t, v18.Words, 23)
	assert.Equal(t, 18, v18.Number)

	firstWord := v18.Words[0]
	assert.Equal(t, "Ὧδε", firstWord.Original)
	assert.Equal(t, "Hōde", firstWord.Transliteration)
	assert.Equal(t, "5602", firstWord.StrongsNumber)
	assert.Equal(t, `Strong's Greek 5602: From an adverb form of hode; in this same spot, i.e. Here or hither.`, firstWord.StrongsText)
	assert.Equal(t, "Adv", firstWord.SyntaxCode)
	assert.Equal(t, "Adverb", firstWord.SyntaxDescription)
	assert.Equal(t, "Here", firstWord.English)

	lastWord := v18.Words[len(v18.Words)-1]
	assert.Equal(t, "ἕξ", lastWord.Original)
	assert.Equal(t, "hex", lastWord.Transliteration)
	assert.Equal(t, "1803", lastWord.StrongsNumber)
	assert.Equal(t, `Strong's Greek 1803: Six. A primary numeral; six.`, lastWord.StrongsText)
	assert.Equal(t, "Adj-NMP", lastWord.SyntaxCode)
	assert.Equal(t, "Adjective - Nominative Masculine Plural", lastWord.SyntaxDescription)
	assert.Equal(t, "six", lastWord.English)
}
