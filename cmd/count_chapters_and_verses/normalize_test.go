package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeIdentifier(t *testing.T) {
	testdata := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
		{
			name:     "Spaces",
			input:    "Song of Songs",
			expected: "Song_of_Songs",
		},
		{
			name:     "Starting with number without space",
			input:    "1Samuel",
			expected: "Samuel1",
		},
		{
			name:     "Starting with number with space",
			input:    "1 Kings",
			expected: "Kings_1",
		},
		{
			name:     "Starting with number followed by chars with space",
			input:    "2nd Chronicles",
			expected: "Chronicles_2nd",
		},
		{
			name:     "Number simple",
			input:    "123",
			expected: "_123",
		},
		{
			name:     "Number with spaces",
			input:    "123 456 789",
			expected: "_123_456_789",
		},
	}
	for _, td := range testdata {
		t.Run(td.name, func(t *testing.T) {
			output := normalizeIdentifier(td.input)
			assert.Equal(t, td.expected, output)
		})
	}
}
