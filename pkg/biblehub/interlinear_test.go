package biblehub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInterlinearChapter_Daniel2(t *testing.T) {
	ch, err := GetInterlinearChapter("daniel", 2)
	require.NoError(t, err)

	assert.Equal(t, "Daniel 2", ch.Title)
	require.Len(t, ch.Verses, 49)

	v49 := ch.Verses[48]
	require.Len(t, v49.Words, 30)
	assert.Equal(t, 49, v49.Number)

	w1 := v49.Words[0]
	assert.Equal(t, "1841", w1.StrongsNumber)
	assert.Equal(t, `Strong's Hebrew 1841: Daniel = God is my judge
 1) the 4th of the greater prophets, taken as hostage in the first deportation to Babylon, because of the gift of God of the interpretation of dreams, he became the 2nd in command of the Babylon empire and lasted through the end of the Babylonian empire and into the Persian empire. His prophecies are the key to the understanding of end time events. Noted for his purity and holiness by contemporary prophet, Ezekiel 
 1a) also, 'Belteshazzar' ( H01095 or H01096)`, w1.StrongsText)
	assert.Equal(t, "wə·ḏā·nî·yêl", w1.Transliteration)
	assert.Equal(t, "וְדָנִיֵּאל֙", w1.Hebrew)
	assert.Equal(t, "And Daniel", w1.English)
}

func TestGetInterlinearChapter_Revelation13(t *testing.T) {
	ch, err := GetInterlinearChapter("revelation", 13)
	require.NoError(t, err)

	assert.Equal(t, "Revelation 13", ch.Title)
}
