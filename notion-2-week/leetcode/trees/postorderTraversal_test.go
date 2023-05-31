package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	tests := map[string]struct {
		input1 *TreeNode
		output []int
	}{
		"1": {
			input1: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			output: []int{1, 2, 3},
		},
		"3": {
			input1: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: nil,
				},
			},
			output: []int{3, 2, 1},
		},
		"2": {
			input1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			output: []int{2, 1},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, postorderTraversal(tt.input1))
		})
	}
}
