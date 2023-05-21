package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := map[string]struct {
		input1 string
		result bool
	}{
		"0": {
			input1: "()",
			result: true,
		},
		"1": {
			input1: "()[]{}",
			result: true,
		},
		"3": {
			input1: "([{}])",
			result: true,
		},
		"2": {
			input1: "(]",
			result: false,
		},
		"4": {
			input1: "((",
			result: false,
		},
		"5": {
			input1: "(",
			result: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.result, isValid(tt.input1))
		})
	}
}
