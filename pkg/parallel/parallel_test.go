package parallel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForI(t *testing.T) {
	t.Run("0-100 simple", func(t *testing.T) {
		result, err := ForI(0, 100, func(i int) (any, error) {
			return i, nil
		})
		require.NoError(t, err)
		assert.Len(t, result, 100)

		assert.Equal(t, 0, result[0])
		assert.Equal(t, 99, result[99])
		assert.Panics(t, func() {
			_ = result[100]
		})
	})
	t.Run("200-300 simple", func(t *testing.T) {
		result, err := ForI(200, 300, func(i int) (any, error) {
			return i, nil
		})
		require.NoError(t, err)
		assert.Len(t, result, 100)

		assert.Equal(t, 200, result[0])
		assert.Equal(t, 299, result[99])
		assert.Panics(t, func() {
			_ = result[100]
		})
	})

	t.Run("10-20 + 3 simple", func(t *testing.T) {
		result, err := ForI(10, 20, func(i int) (any, error) {
			return i + 3, nil
		})
		require.NoError(t, err)
		assert.Len(t, result, 10)

		assert.Equal(t, 10+3, result[0])
		assert.Equal(t, 19+3, result[9])
		assert.Panics(t, func() {
			_ = result[10]
		})
	})
}
