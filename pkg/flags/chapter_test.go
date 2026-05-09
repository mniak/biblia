package flags

import (
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChapterFlag_Set(t *testing.T) {
	aliases := map[string]bible.BookCode{
		"Genesis": bible.Genesis,
		"Gn":      bible.Genesis,
		"IJoao":   bible.IJohn,
	}

	tests := []struct {
		name    string
		input   string
		want    bible.ChapterRef
		wantErr bool
	}{
		{
			name:  "by UBS code",
			input: "GEN1",
			want:  bible.ChapterRef{Book: bible.Genesis, Chapter: 1},
		},
		{
			name:  "by alias",
			input: "Genesis1",
			want:  bible.ChapterRef{Book: bible.Genesis, Chapter: 1},
		},
		{
			name:  "with roman numeral",
			input: "IJoao3",
			want:  bible.ChapterRef{Book: bible.IJohn, Chapter: 3},
		},
		{
			name:    "invalid",
			input:   "Unknown1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := NewChapterFlag(aliases)
			err := flag.Set(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, flag.Value())
		})
	}
}

func TestChapterFlag_String(t *testing.T) {
	aliases := map[string]bible.BookCode{"Genesis": bible.Genesis}

	t.Run("empty when not set", func(t *testing.T) {
		flag := NewChapterFlag(aliases)
		assert.Equal(t, "", flag.String())
	})

	t.Run("returns formatted ref when set", func(t *testing.T) {
		flag := NewChapterFlag(aliases)
		_ = flag.Set("Genesis1")
		assert.Equal(t, "GEN1", flag.String())
	})
}

func TestChapterFlag_Type(t *testing.T) {
	flag := NewChapterFlag(nil)
	assert.Equal(t, "chapter", flag.Type())
}

func TestChapterFlag_IsSet(t *testing.T) {
	aliases := map[string]bible.BookCode{"Genesis": bible.Genesis}

	t.Run("false when not set", func(t *testing.T) {
		flag := NewChapterFlag(aliases)
		assert.False(t, flag.IsSet())
	})

	t.Run("true when set", func(t *testing.T) {
		flag := NewChapterFlag(aliases)
		_ = flag.Set("Genesis1")
		assert.True(t, flag.IsSet())
	})
}
