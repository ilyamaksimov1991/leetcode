package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	tests := map[string]struct {
		input1 string
		result int
	}{
		"0": {
			input1: "Hello World",
			result: 5,
		},
		"1": {
			input1: "   fly me   to   the moon  ",
			result: 4,
		},
		"2": {
			input1: "luffy is still joyboy",
			result: 6,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.result, lengthOfLastWord(tt.input1))
		})
	}
}
