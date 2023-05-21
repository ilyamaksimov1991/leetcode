package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := map[string]struct {
		input  []int
		result []int
	}{
		"0": {
			input:  []int{0, 1, 0, 3, 12},
			result: []int{1, 3, 12, 0, 0},
		},
		"01": {
			input:  []int{1, 2, 3, 0},
			result: []int{1, 2, 3, 0},
		},
		"1": {
			input:  []int{0},
			result: []int{0},
		},
		"2": {
			input:  []int{0, 0, 1},
			result: []int{1, 0, 0},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.result, moveZeroes2(tt.input))
		})
	}
}
