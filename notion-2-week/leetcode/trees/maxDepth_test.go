package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := map[string]struct {
		input1 *TreeNode
		output int
	}{
		"1": {
			input1: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			output: 3,
		},
		"2": {
			input1: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			output: 5,
		},
		"3": {
			input1: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			output: 2,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, maxDepth(tt.input1))
		})
	}
}
