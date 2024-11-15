package runeutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	_ RuneWalker = NewRuneWalker("")
	_ RuneWalker = NewReverseRuneWalker("")
	_ RuneWalker = NewRuneWalkerFromRunes([]rune{})
)

func TestRuneWalker(t *testing.T) {
	word := "1234"
	sut := NewRuneWalker(word)

	// sut.Rune = 0
	require.True(t, sut.Walk())
	require.Equal(t, '1', sut.Rune())

	require.False(t, sut.WalkBack())

	require.True(t, sut.Walk())
	require.Equal(t, '2', sut.Rune())

	require.True(t, sut.WalkBack())
	require.Equal(t, '1', sut.Rune())

	require.False(t, sut.WalkBack())
	require.False(t, sut.WalkBack())
	require.Equal(t, '1', sut.Rune())

	require.True(t, sut.Walk())
	require.Equal(t, '2', sut.Rune())
	require.True(t, sut.Walk())
	require.Equal(t, '3', sut.Rune())
	require.True(t, sut.Walk())
	require.Equal(t, '4', sut.Rune())

	require.False(t, sut.Walk())
	require.False(t, sut.Walk())
}

func TestReverseRuneWalker(t *testing.T) {
	word := "4321"
	sut := NewReverseRuneWalker(word)

	// sut.Rune = 0
	require.True(t, sut.Walk())
	require.Equal(t, '1', sut.Rune())

	require.False(t, sut.WalkBack())

	require.True(t, sut.Walk())
	require.Equal(t, '2', sut.Rune())

	require.True(t, sut.WalkBack())
	require.Equal(t, '1', sut.Rune())

	require.False(t, sut.WalkBack())
	require.False(t, sut.WalkBack())
	require.Equal(t, '1', sut.Rune())

	require.True(t, sut.Walk())
	require.Equal(t, '2', sut.Rune())
	require.True(t, sut.Walk())
	require.Equal(t, '3', sut.Rune())
	require.True(t, sut.Walk())
	require.Equal(t, '4', sut.Rune())

	require.False(t, sut.Walk())
	require.False(t, sut.Walk())
}
