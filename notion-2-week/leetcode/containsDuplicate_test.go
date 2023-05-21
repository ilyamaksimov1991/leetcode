package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	tests := map[string]struct {
		input  []int
		output bool
	}{
		"0": {
			input:  []int{1, 2, 3, 1},
			output: true,
		},
		"1": {
			input:  []int{1, 2, 3, 4},
			output: false,
		},
		"2": {
			input:  []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			output: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, containsDuplicate(tt.input))
		})
	}
}
