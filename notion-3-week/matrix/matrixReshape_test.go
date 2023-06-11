package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	tests := map[string]struct {
		input1 [][]int
		c      int
		r      int
		output [][]int
	}{
		"1": {
			input1: [][]int{
				{1, 2},
				{3, 4},
			},
			c: 4,
			r: 1,
			output: [][]int{
				{1, 2, 3, 4},
			},
		},
		"2": {
			input1: [][]int{
				{1, 2},
				{3, 4},
			},
			c: 4,
			r: 2,
			output: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, matrixReshape(tt.input1, tt.r, tt.c))
		})
	}
}
