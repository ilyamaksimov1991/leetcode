package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	tests := map[string]struct {
		input1 *TreeNode
		input2 *TreeNode
		output bool
	}{
		"1": {
			input1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			input2: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			output: true,
		},
		"2": {
			input1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			input2: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			output: false,
		},
		"3": {
			input1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			input2: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			output: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, isSameTree(tt.input1, tt.input2))
		})
	}
}
