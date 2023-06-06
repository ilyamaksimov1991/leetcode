package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := map[string]struct {
		input1 []int
		m1     int
		input2 []int
		m2     int
		output []int
	}{
		"0": {
			input1: []int{1, 2, 3, 0, 0, 0},
			m1:     3,
			input2: []int{2, 5, 6},
			m2:     3,
			output: []int{1, 2, 2, 3, 5, 6},
		},
		"1": {
			input1: []int{1},
			m1:     1,
			input2: []int{},
			m2:     0,
			output: []int{1},
		},
		"2": {
			input1: []int{0},
			m1:     0,
			input2: []int{1},
			m2:     1,
			output: []int{1},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, merge(tt.input1, tt.m1, tt.input2, tt.m2))
		})
	}
}
