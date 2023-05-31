package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := map[string]struct {
		input  []int
		output []int
		k      int
	}{
		"0": {
			input:  []int{1, 1, 2},
			output: []int{1, 2},
			k:      2,
		},
		"1": {
			input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output: []int{0, 1, 2, 3, 4},
			k:      5,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			o, k := removeDuplicates(tt.input)
			assert.Equal(t, tt.output, o)
			assert.Equal(t, tt.k, k)
		})
	}
}
