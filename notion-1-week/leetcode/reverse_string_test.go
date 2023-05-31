package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseString(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected []string
	}{
		"0": {
			input:    []string{"d", "e", "s"},
			expected: []string{"s", "e", "d"},
		},
		"1": {
			input:    []string{"H", "e", "e", "h"},
			expected: []string{"h", "e", "e", "H"},
		},
		"one": {
			input:    []string{"h", "e", "l", "l", "o"},
			expected: []string{"o", "l", "l", "e", "h"},
		},
		"two": {
			input:    []string{"H", "a", "n", "n", "a", "h"},
			expected: []string{"h", "a", "n", "n", "a", "H"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, reverseString(tt.input))
		})
	}
}
