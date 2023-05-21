package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	tests := map[string]struct {
		input1 *TreeNode
		output []int
	}{
		"1": {
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
			output: []int{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, preorderTraversal(tt.input1))
		})
	}
}
