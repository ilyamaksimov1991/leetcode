package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchBST(t *testing.T) {
	tests := map[string]struct {
		input1 *TreeNode
		val    int
		output *TreeNode
	}{
		"1": {
			input1: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			val: 2,
			output: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		"2": {
			input1: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			val:    5,
			output: nil,
		},
		"3": {
			input1: &TreeNode{
				Val: 18,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 22,
					Right: &TreeNode{
						Val: 63,
						Right: &TreeNode{
							Val: 84,
						},
					},
				},
			},
			val: 63,
			output: &TreeNode{
				Val: 63,
				Right: &TreeNode{
					Val: 84,
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, searchBST(tt.input1, tt.val))
		})
	}
}
