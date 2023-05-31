package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := map[string]struct {
		input1 []string
		result string
	}{
		"0": {
			input1: []string{"flower", "flow", "flight"},
			result: "fl",
		},
		"2": {
			input1: []string{"dog", "racecar", "car"},
			result: "",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.result, longestCommonPrefix(tt.input1))
		})
	}
}
