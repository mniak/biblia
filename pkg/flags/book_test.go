package flags

import (
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookFlag_Set(t *testing.T) {
	aliases := map[string]bible.BookCode{
		"Gn":      bible.Genesis,
		"Gen":     bible.Genesis,
		"Genesis": bible.Genesis,
		"Ex":      bible.Exodus,
		"Exod":    bible.Exodus,
		"Exodus":  bible.Exodus,
		"Mt":      bible.Matthew,
		"Matt":    bible.Matthew,
		"Matthew": bible.Matthew,
	}

	tests := []struct {
		name     string
		input    string
		wantCode bible.BookCode
		wantErr  bool
	}{
		{name: "by code uppercase", input: "GEN", wantCode: bible.Genesis},
		{name: "by code lowercase", input: "gen", wantCode: bible.Genesis},
		{name: "by name", input: "Genesis", wantCode: bible.Genesis},
		{name: "by name case insensitive", input: "GENESIS", wantCode: bible.Genesis},
		{name: "by alias", input: "Gn", wantCode: bible.Genesis},
		{name: "by alias case insensitive", input: "gn", wantCode: bible.Genesis},
		{name: "by another alias", input: "Mt", wantCode: bible.Matthew},
		{name: "unknown book", input: "Unknown", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := NewBookFlag(aliases)
			err := flag.Set(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantCode, flag.Value())
		})
	}
}

func TestBookFlag_String(t *testing.T) {
	aliases := map[string]bible.BookCode{
		"Genesis": bible.Genesis,
	}

	t.Run("empty when not set", func(t *testing.T) {
		flag := NewBookFlag(aliases)
		assert.Equal(t, "", flag.String())
	})

	t.Run("returns code when set", func(t *testing.T) {
		flag := NewBookFlag(aliases)
		_ = flag.Set("Genesis")
		assert.Equal(t, "GEN", flag.String())
	})
}

func TestBookFlag_Type(t *testing.T) {
	flag := NewBookFlag(nil)
	assert.Equal(t, "book", flag.Type())
}

func TestBookFlag_IsSet(t *testing.T) {
	aliases := map[string]bible.BookCode{
		"Genesis": bible.Genesis,
	}

	t.Run("false when not set", func(t *testing.T) {
		flag := NewBookFlag(aliases)
		assert.False(t, flag.IsSet())
	})

	t.Run("true when set", func(t *testing.T) {
		flag := NewBookFlag(aliases)
		_ = flag.Set("Genesis")
		assert.True(t, flag.IsSet())
	})
}
