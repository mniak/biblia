package flags

import (
	"testing"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookIndexRangeFlag_Set(t *testing.T) {
	aliases := map[string]bible.BookCode{
		"Gn":      bible.Genesis,
		"Gen":     bible.Genesis,
		"Genesis": bible.Genesis,
		"Ex":      bible.Exodus,
		"Mt":      bible.Matthew,
		"Matt":    bible.Matthew,
	}

	tests := []struct {
		name    string
		input   string
		want    []bible.BookIndex
		wantErr bool
	}{
		{
			name:  "single by code",
			input: "GEN",
			want:  []bible.BookIndex{1},
		},
		{
			name:  "single by code lowercase",
			input: "gen",
			want:  []bible.BookIndex{1},
		},
		{
			name:  "single by alias",
			input: "Genesis",
			want:  []bible.BookIndex{1},
		},
		{
			name:  "range by code",
			input: "GEN-DEU",
			want:  []bible.BookIndex{1, 2, 3, 4, 5},
		},
		{
			name:  "multiple single books",
			input: "GEN,LEV,DEU",
			want:  []bible.BookIndex{1, 3, 5},
		},
		{
			name:  "gospels by code",
			input: "MAT-JHN",
			want:  []bible.BookIndex{40, 41, 42, 43},
		},
		{
			name:  "combined ranges",
			input: "GEN-LEV,MAT-LUK",
			want:  []bible.BookIndex{1, 2, 3, 40, 41, 42},
		},
		{
			name:    "unknown book",
			input:   "UNKNOWN",
			wantErr: true,
		},
		{
			name:    "number not allowed",
			input:   "1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := NewBookIndexRangeFlag(aliases)
			err := flag.Set(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, flag.Collect())
		})
	}
}

func TestBookIndexRangeFlag_Values(t *testing.T) {
	flag := NewBookIndexRangeFlag(nil)
	_ = flag.Set("GEN-LEV,MAT")

	var indices []bible.BookIndex
	for idx := range flag.Values() {
		indices = append(indices, idx)
	}

	assert.Equal(t, []bible.BookIndex{1, 2, 3, 40}, indices)
}

func TestBookIndexRangeFlag_Contains(t *testing.T) {
	flag := NewBookIndexRangeFlag(nil)
	_ = flag.Set("GEN-DEU,MAT-JHN")

	assert.True(t, flag.Contains(1))  // Genesis
	assert.True(t, flag.Contains(5))  // Deuteronomy
	assert.True(t, flag.Contains(40)) // Matthew
	assert.True(t, flag.Contains(43)) // John
	assert.False(t, flag.Contains(6)) // Joshua
	assert.False(t, flag.Contains(39))
	assert.False(t, flag.Contains(66)) // Revelation
}

func TestBookIndexRangeFlag_IterateAsBookCode(t *testing.T) {
	flag := NewBookIndexRangeFlag(nil)
	_ = flag.Set("GEN-LEV")

	var codes []bible.BookCode
	for idx := range flag.Values() {
		codes = append(codes, idx.BookCode())
	}

	assert.Equal(t, []bible.BookCode{bible.Genesis, bible.Exodus, bible.Leviticus}, codes)
}

func TestBookIndexRangeFlag_Type(t *testing.T) {
	flag := NewBookIndexRangeFlag(nil)
	assert.Equal(t, "range", flag.Type())
}

func TestBookIndexRangeFlag_IsSet(t *testing.T) {
	t.Run("false when not set", func(t *testing.T) {
		flag := NewBookIndexRangeFlag(nil)
		assert.False(t, flag.IsSet())
	})

	t.Run("true when set", func(t *testing.T) {
		flag := NewBookIndexRangeFlag(nil)
		_ = flag.Set("GEN-DEU")
		assert.True(t, flag.IsSet())
	})
}
