package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPositiveIntFlag_Set(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "valid positive integer", input: "42", want: 42},
		{name: "valid one", input: "1", want: 1},
		{name: "large number", input: "1000000", want: 1000000},
		{name: "zero is invalid", input: "0", wantErr: true},
		{name: "negative is invalid", input: "-5", wantErr: true},
		{name: "non-integer", input: "abc", wantErr: true},
		{name: "float", input: "3.14", wantErr: true},
		{name: "empty", input: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := NewPositiveIntFlag(0)
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

func TestPositiveIntFlag_String(t *testing.T) {
	t.Run("returns default value as string", func(t *testing.T) {
		flag := NewPositiveIntFlag(10)
		assert.Equal(t, "10", flag.String())
	})

	t.Run("returns set value as string", func(t *testing.T) {
		flag := NewPositiveIntFlag(10)
		_ = flag.Set("25")
		assert.Equal(t, "25", flag.String())
	})
}

func TestPositiveIntFlag_Type(t *testing.T) {
	flag := NewPositiveIntFlag(0)
	assert.Equal(t, "positive-int", flag.Type())
}

func TestPositiveIntFlag_IsSet(t *testing.T) {
	t.Run("false when not set", func(t *testing.T) {
		flag := NewPositiveIntFlag(10)
		assert.False(t, flag.IsSet())
	})

	t.Run("true when set", func(t *testing.T) {
		flag := NewPositiveIntFlag(10)
		_ = flag.Set("25")
		assert.True(t, flag.IsSet())
	})
}
