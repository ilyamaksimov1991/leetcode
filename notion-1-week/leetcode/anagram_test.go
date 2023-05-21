package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	tests := map[string]struct {
		input1 string
		input2 string
		result bool
	}{
		"0": {
			input1: "anagram",
			input2: "nagaram",
			result: true,
		},
		"1": {
			input1: "rat",
			input2: "car",
			result: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.result, isAnagram(tt.input1, tt.input2))
		})
	}
}
