package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	tests := map[string]struct {
		input1 *TreeNode
		output []int
	}{
		"1": {
			input1: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			output: []int{1, 3, 2},
		},
		"2": {
			input1: nil,
			output: nil,
		},
		"3": {
			input1: &TreeNode{
				Val: 1,
			},
			output: []int{1},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, inorderTraversal(tt.input1))
		})
	}
}
