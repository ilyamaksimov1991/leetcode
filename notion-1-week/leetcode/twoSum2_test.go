package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum2(t *testing.T) {
	tests := map[string]struct {
		nums   []int
		target int
		output []int
	}{
		"0": {
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{1, 2},
		},
		"1": {
			nums:   []int{2, 3, 4},
			target: 6,
			output: []int{1, 3},
		},
		"2": {
			nums:   []int{-1, 0},
			target: -1,
			output: []int{1, 2},
		},
		"3": {
			nums:   []int{1, 2, 3, 4, 4, 9, 56, 90},
			target: 8,
			output: []int{4, 5},
		},
		//"4": {
		//	nums:   []int{-3, 4, 3, 90},
		//	target: 0,
		//	output: []int{0, 2},
		//},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, twoSum2(tt.nums, tt.target))
		})
	}
}
