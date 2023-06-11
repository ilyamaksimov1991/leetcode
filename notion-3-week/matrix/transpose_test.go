package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_transpose(t *testing.T) {
	tests := map[string]struct {
		input1 [][]int
		output [][]int
	}{
		"1": {
			input1: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			output: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		"2": {
			input1: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			output: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, transpose(tt.input1))
		})
	}
}
