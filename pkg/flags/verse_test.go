package flags

import (
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerseFlag_Set(t *testing.T) {
	aliases := map[string]bible.BookCode{
		"Genesis": bible.Genesis,
		"Gn":      bible.Genesis,
		"IJoao":   bible.IJohn,
		"ijoao":   bible.IJohn,
	}

	tests := []struct {
		name    string
		input   string
		want    bible.VerseRef
		wantErr bool
	}{
		{
			name:  "by UBS code",
			input: "GEN1:1",
			want:  bible.VerseRef{Book: bible.Genesis, Chapter: 1, Verse: 1},
		},
		{
			name:  "by alias",
			input: "Genesis1:1",
			want:  bible.VerseRef{Book: bible.Genesis, Chapter: 1, Verse: 1},
		},
		{
			name:  "with roman numeral",
			input: "IJoao4:8",
			want:  bible.VerseRef{Book: bible.IJohn, Chapter: 4, Verse: 8},
		},
		{
			name:  "with roman numeral lowercase",
			input: "ijoao4:8",
			want:  bible.VerseRef{Book: bible.IJohn, Chapter: 4, Verse: 8},
		},
		{
			name:    "missing colon",
			input:   "Genesis1",
			wantErr: true,
		},
		{
			name:    "invalid book",
			input:   "Unknown1:1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := NewVerseFlag(aliases)
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

func TestVerseFlag_String(t *testing.T) {
	aliases := map[string]bible.BookCode{"Genesis": bible.Genesis}

	t.Run("empty when not set", func(t *testing.T) {
		flag := NewVerseFlag(aliases)
		assert.Equal(t, "", flag.String())
	})

	t.Run("returns formatted ref when set", func(t *testing.T) {
		flag := NewVerseFlag(aliases)
		_ = flag.Set("Genesis1:1")
		assert.Equal(t, "GEN1:1", flag.String())
	})
}

func TestVerseFlag_Type(t *testing.T) {
	flag := NewVerseFlag(nil)
	assert.Equal(t, "verse", flag.Type())
}

func TestVerseFlag_IsSet(t *testing.T) {
	aliases := map[string]bible.BookCode{"Genesis": bible.Genesis}

	t.Run("false when not set", func(t *testing.T) {
		flag := NewVerseFlag(aliases)
		assert.False(t, flag.IsSet())
	})

	t.Run("true when set", func(t *testing.T) {
		flag := NewVerseFlag(aliases)
		_ = flag.Set("Genesis1:1")
		assert.True(t, flag.IsSet())
	})
}
