package hashTables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	tests := map[string]struct {
		input  []int
		output int
	}{
		"1": {
			input:  []int{3, 2, 3},
			output: 3,
		},
		"2": {
			input:  []int{2, 2, 1, 1, 1, 2, 2},
			output: 2,
		},
		"3": {
			input:  []int{1},
			output: 1,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, majorityElement(tt.input))
		})
	}
}
