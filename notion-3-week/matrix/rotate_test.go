package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotate(t *testing.T) {
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
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			rotate(tt.input1)
			assert.Equal(t, tt.output, tt.input1)
		})
	}
}
