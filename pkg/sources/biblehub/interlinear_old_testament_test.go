package biblehub

import (
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ bible.TestamentLoader = interlinearOTLoader{}

func TestLoadOTChapter(t *testing.T) {
	chapter, err := loadOTChapter("genesis", 1)
	require.NoError(t, err)

	assert.NotEmpty(t, chapter)
	assert.Equal(t, 1, chapter.Number)
	assert.Len(t, chapter.Verses, 31)
}
