package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tests := map[string]struct {
		input  []int
		result []int
	}{
		"0": {
			input:  []int{-4, -1, 0, 3, 10},
			result: []int{0, 1, 9, 16, 100},
		},
		"1": {
			input:  []int{-7, -3, 2, 3, 11},
			result: []int{4, 9, 9, 49, 121},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.result, sortedSquares(tt.input))
		})
	}
}
