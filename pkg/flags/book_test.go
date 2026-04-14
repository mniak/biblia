package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookFlag_Set(t *testing.T) {
	books := []BookInfo{
		{Index: 0, Code: "GEN", Name: "Genesis", Aliases: []string{"Gn", "Gen"}, Testament: OldTestament},
		{Index: 1, Code: "EXO", Name: "Exodus", Aliases: []string{"Ex", "Exod"}, Testament: OldTestament},
		{Index: 39, Code: "MAT", Name: "Matthew", Aliases: []string{"Mt", "Matt"}, Testament: NewTestament},
	}

	tests := []struct {
		name     string
		input    string
		wantCode string
		wantErr  bool
	}{
		{name: "by code uppercase", input: "GEN", wantCode: "GEN"},
		{name: "by code lowercase", input: "gen", wantCode: "GEN"},
		{name: "by name", input: "Genesis", wantCode: "GEN"},
		{name: "by name case insensitive", input: "GENESIS", wantCode: "GEN"},
		{name: "by alias", input: "Gn", wantCode: "GEN"},
		{name: "by alias case insensitive", input: "gn", wantCode: "GEN"},
		{name: "by another alias", input: "Mt", wantCode: "MAT"},
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
		{Index: 0, Code: "GEN", Name: "Genesis", Aliases: []string{"Gn"}, Testament: OldTestament},
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
		{Index: 0, Code: "GEN", Name: "Genesis", Testament: OldTestament},
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
