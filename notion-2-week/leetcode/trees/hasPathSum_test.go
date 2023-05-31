package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	tests := map[string]struct {
		input1 *TreeNode
		val    int
		output bool
	}{
		"1": {
			input1: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			val:    22,
			output: true,
		},
		"2": {
			input1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			val:    5,
			output: false,
		},
		"3": {
			input1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			val:    1,
			output: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, hasPathSum(tt.input1, tt.val))
		})
	}
}
