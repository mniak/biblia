package flags

import (
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookFlag_Set(t *testing.T) {
	books := []BookInfo{
		{Code: bible.Genesis, Aliases: []string{"Gn", "Gen", "Genesis"}},
		{Code: bible.Exodus, Aliases: []string{"Ex", "Exod", "Exodus"}},
		{Code: bible.Matthew, Aliases: []string{"Mt", "Matt", "Matthew"}},
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
			flag := NewBookFlag(books)
			err := flag.Set(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantCode, flag.Value().Code)
		})
	}
}

func TestBookFlag_String(t *testing.T) {
	books := []BookInfo{
		{Code: bible.Genesis, Aliases: []string{"Gn", "Genesis"}},
	}

	t.Run("empty when not set", func(t *testing.T) {
		flag := NewBookFlag(books)
		assert.Equal(t, "", flag.String())
	})

	t.Run("returns code when set", func(t *testing.T) {
		flag := NewBookFlag(books)
		_ = flag.Set("Genesis")
		assert.Equal(t, "GEN", flag.String())
	})
}

func TestBookFlag_Type(t *testing.T) {
	flag := NewBookFlag(nil)
	assert.Equal(t, "book", flag.Type())
}

func TestBookFlag_IsSet(t *testing.T) {
	books := []BookInfo{
		{Code: bible.Genesis, Aliases: []string{"Genesis"}},
	}

	t.Run("false when not set", func(t *testing.T) {
		flag := NewBookFlag(books)
		assert.False(t, flag.IsSet())
	})

	t.Run("true when set", func(t *testing.T) {
		flag := NewBookFlag(books)
		_ = flag.Set("Genesis")
		assert.True(t, flag.IsSet())
	})
}
