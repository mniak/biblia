package biblehub

import (
	"testing"

	"github.com/mniak/biblia/internal/test"
	"github.com/mniak/biblia/pkg/biblehub/biblehubtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInterlinearChapter_Daniel2(t *testing.T) {
	ex := _Extractor{
		Downloader: biblehubtest.FakeDownloader,
	}

	ch, err := ex.GetInterlinearChapter("daniel", 2)
	require.NoError(t, err)
	assert.Equal(t, Hebrew, ch.Language)

	assert.Equal(t, "Daniel 2", ch.Title)
	require.Len(t, ch.Verses, 49)

	v49 := ch.Verses[48]
	require.Len(t, v49.Words, 30)
	assert.Equal(t, 49, v49.Number)

	firstWord := v49.Words[0]
	assert.Equal(t, "1841", firstWord.StrongsNumber)
	test.AssertEqualTrim(t,
		`Strong's Hebrew 1841: Daniel = God is my judge
		1) the 4th of the greater prophets, taken as hostage in the first deportation to Babylon, because of the gift of God of the interpretation of dreams, he became the 2nd in command of the Babylon empire and lasted through the end of the Babylonian empire and into the Persian empire. His prophecies are the key to the understanding of end time events. Noted for his purity and holiness by contemporary prophet, Ezekiel 
		1a) also, 'Belteshazzar' ( H01095 or H01096)`,
		firstWord.StrongsText)
	assert.Equal(t, "wə·ḏā·nî·yêl", firstWord.Transliteration)
	assert.Equal(t, "וְדָנִיֵּאל֙", firstWord.Original)
	assert.Equal(t, "And Daniel", firstWord.English)

	previousToLastWord := v49.Words[len(v49.Words)-2]
	assert.Equal(t, "4430", previousToLastWord.StrongsNumber)
	test.AssertEqualTrim(t, `Strong's Hebrew 4430: 1) king`, previousToLastWord.StrongsText)
	assert.Equal(t, "mal·kā.", previousToLastWord.Transliteration)
	assert.Equal(t, "מַלְכָּֽא׃", previousToLastWord.Original)
	assert.Equal(t, "of king the", previousToLastWord.English)
}

func TestGetInterlinearChapter_Revelation13(t *testing.T) {
	ex := _Extractor{
		Downloader: biblehubtest.FakeDownloader,
	}

	ch, err := ex.GetInterlinearChapter("revelation", 13)
	require.NoError(t, err)
	assert.Equal(t, Greek, ch.Language)

	assert.Equal(t, "Revelation 13", ch.Title)
	require.Len(t, ch.Verses, 18)

	v18 := ch.Verses[17]
	require.Len(t, v18.Words, 23)
	assert.Equal(t, 18, v18.Number)

	firstWord := v18.Words[0]
	assert.Equal(t, "5602", firstWord.StrongsNumber)
	assert.Equal(t, `Strong's Greek 5602: From an adverb form of hode; in this same spot, i.e. Here or hither.`, firstWord.StrongsText)
	assert.Equal(t, "Hōde", firstWord.Transliteration)
	assert.Equal(t, "Ὧδε", firstWord.Original)
	assert.Equal(t, "Here", firstWord.English)

	lastWord := v18.Words[len(v18.Words)-1]
	assert.Equal(t, "1803", lastWord.StrongsNumber)
	assert.Equal(t, `Strong's Greek 1803: Six. A primary numeral; six.`, lastWord.StrongsText)
	assert.Equal(t, "hex", lastWord.Transliteration)
	assert.Equal(t, "ἕξ", lastWord.Original)
	assert.Equal(t, "six", lastWord.English)
}
