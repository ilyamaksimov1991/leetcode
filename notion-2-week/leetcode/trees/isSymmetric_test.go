package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	tests := map[string]struct {
		input1 *TreeNode
		output bool
	}{
		"1": {
			input1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			output: true,
		},
		"2": {
			input1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			output: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, isSymmetric(tt.input1))
		})
	}
}
