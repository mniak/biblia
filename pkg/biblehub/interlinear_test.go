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
	assert.Len(t, ch.Verses, 49)
}

func TestGetInterlinearChapter_Revelation13(t *testing.T) {
	ch, err := GetInterlinearChapter("revelation", 13)
	require.NoError(t, err)

	assert.Equal(t, "Revelation 13", ch.Title)
}
