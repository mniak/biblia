package biblehub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetInterlinearChapter(t *testing.T) {
	ch, err := GetInterlinearChapter("genesis", 1)
	require.NoError(t, err)

	assert.Equal(t, "Genesis 1", ch.Title)
}
