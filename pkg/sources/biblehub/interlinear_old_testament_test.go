package biblehub

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadOTChapter(t *testing.T) {
	book, err := loadOTChapter("genesis", 1)
	require.NoError(t, err)

	assert.NotEmpty(t, book)
}
