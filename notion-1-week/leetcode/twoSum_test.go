package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		output []int
	}{
		"0": {
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		"1": {
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		"2": {
			nums:   []int{3, 3},
			target: 6,
			output: []int{0, 1},
		},
		"3": {
			nums:   []int{0, 4, 3, 0},
			target: 0,
			output: []int{0, 3},
		},
		"4": {
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			output: []int{0, 2},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, twoSum(tt.nums, tt.target))
		})
	}
}
