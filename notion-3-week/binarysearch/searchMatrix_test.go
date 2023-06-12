package binarysearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	tests := map[string]struct {
		input1 [][]int
		target int
		output bool
	}{
		"1": {
			input1: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			output: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, searchMatrix(tt.input1, tt.target))
		})
	}
}
