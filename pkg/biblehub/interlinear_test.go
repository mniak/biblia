package biblehub

import (
	"fmt"
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

	w1 := v49.Words[0]
	assert.Equal(t, "1841", w1.StrongsNumber)
	assert.Equal(t, "wə·ḏā·nî·yêl", w1.Transliteration)
	assert.Equal(t, "וְדָנִיֵּאל֙", w1.Hebrew)
	assert.Equal(t, "And Daniel", w1.English)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(w1)
}

func TestGetInterlinearChapter_Revelation13(t *testing.T) {
	ch, err := GetInterlinearChapter("revelation", 13)
	require.NoError(t, err)

	assert.Equal(t, "Revelation 13", ch.Title)
}
