package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pathSum(t *testing.T) {
	tests := map[string]struct {
		input1    *TreeNode
		targetSum int
		output    [][]int
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
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 22,
			output:    [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.output, pathSum(tt.input1, tt.targetSum))
		})
	}
}
