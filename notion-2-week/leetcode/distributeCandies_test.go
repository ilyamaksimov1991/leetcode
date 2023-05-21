package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	tests := map[string]struct {
		input  []int
		output int
	}{
		"0": {
			input:  []int{1, 1, 2, 2, 3, 3},
			output: 3,
		},
		"1": {
			input:  []int{1, 1, 2, 3},
			output: 2,
		},
		"2": {
			input:  []int{6, 6, 6, 6},
			output: 1,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, distributeCandies(tt.input))
		})
	}
}
