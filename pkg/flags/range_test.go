package flags

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRangeFlag_Set(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{name: "single value", input: "5", want: []int{5}},
		{name: "simple range", input: "2-7", want: []int{2, 3, 4, 5, 6, 7}},
		{name: "multiple single values", input: "2,5,7", want: []int{2, 5, 7}},
		{name: "mixed ranges and values", input: "1-3,11-14,20", want: []int{1, 2, 3, 11, 12, 13, 14, 20}},
		{name: "single value range", input: "5-5", want: []int{5}},
		{name: "with spaces", input: " 1 , 3 - 5 ", want: []int{1, 3, 4, 5}},
		{name: "invalid range order", input: "7-2", wantErr: true},
		{name: "invalid start", input: "abc-5", wantErr: true},
		{name: "invalid end", input: "5-abc", wantErr: true},
		{name: "invalid single", input: "abc", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := NewIntRangeFlag()
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

func TestRangeFlag_String(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "single value", input: "5", want: "5"},
		{name: "range", input: "2-7", want: "2-7"},
		{name: "multiple", input: "1,3-5,7", want: "1,3-5,7"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := NewIntRangeFlag()
			if tt.input != "" {
				_ = flag.Set(tt.input)
			}
			assert.Equal(t, tt.want, flag.String())
		})
	}
}

func TestRangeFlag_Type(t *testing.T) {
	flag := NewIntRangeFlag()
	assert.Equal(t, "range", flag.Type())
}

func TestRangeFlag_Values(t *testing.T) {
	flag := NewIntRangeFlag()
	_ = flag.Set("1-3,10")

	var values []int
	for v := range flag.Values() {
		values = append(values, v)
	}

	assert.Equal(t, []int{1, 2, 3, 10}, values)
}

func TestRangeFlag_Values_EarlyBreak(t *testing.T) {
	flag := NewIntRangeFlag()
	_ = flag.Set("1-100")

	var values []int
	for v := range flag.Values() {
		values = append(values, v)
		if v == 5 {
			break
		}
	}

	assert.Equal(t, []int{1, 2, 3, 4, 5}, values)
}

func TestRangeFlag_Contains(t *testing.T) {
	flag := NewIntRangeFlag()
	_ = flag.Set("1-3,10-12")

	tests := []struct {
		value int
		want  bool
	}{
		{value: 0, want: false},
		{value: 1, want: true},
		{value: 2, want: true},
		{value: 3, want: true},
		{value: 4, want: false},
		{value: 9, want: false},
		{value: 10, want: true},
		{value: 11, want: true},
		{value: 12, want: true},
		{value: 13, want: false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, flag.Contains(tt.value), "Contains(%d)", tt.value)
	}
}

func TestRangeFlag_IsSet(t *testing.T) {
	t.Run("false when not set", func(t *testing.T) {
		flag := NewIntRangeFlag()
		assert.False(t, flag.IsSet())
	})

	t.Run("true when set", func(t *testing.T) {
		flag := NewIntRangeFlag()
		_ = flag.Set("1-5")
		assert.True(t, flag.IsSet())
	})
}

func TestRangeFlag_CustomType(t *testing.T) {
	type ChapterNum int

	flag := NewRangeFlag(func(s string) (ChapterNum, error) {
		n, err := parseInt(s)
		return ChapterNum(n), err
	})

	_ = flag.Set("1-3")
	values := flag.Collect()

	assert.Equal(t, []ChapterNum{1, 2, 3}, values)
}

func parseInt(s string) (int, error) {
	var result int
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid character: %c", c)
		}
		result = result*10 + int(c-'0')
	}
	return result, nil
}
