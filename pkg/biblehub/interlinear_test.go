package biblehub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInterlinearChapter(t *testing.T) {
	t.Run("Daniel 2", func(t *testing.T) {
		ch, err := GetInterlinearChapter("daniel", 2)
		require.NoError(t, err)

		assert.Equal(t, "Daniel 2", ch.Title)
	})

	t.Run("Revelation 13", func(t *testing.T) {
		ch, err := GetInterlinearChapter("revelation", 13)
		require.NoError(t, err)

		assert.Equal(t, "Revelation 13", ch.Title)
	})
}
