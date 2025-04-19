package test

import (
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func AssertEqualTrim(t *testing.T, expected, actual string) {
	expected = trimMultiline(expected)
	actual = trimMultiline(actual)
	assert.Equal(t, expected, actual)
}

func trimMultiline(text string) string {
	lines := strings.Split(text, "\n")
	lines = lo.Map[string, string](lines, func(s string, i int) string {
		s = strings.TrimSpace(s)
		return s
	})
	return strings.Join(lines, "\n")
}
