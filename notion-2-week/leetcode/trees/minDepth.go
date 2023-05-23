package trees

// time O(n)
// memory O(1)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return 1
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	return min(r, l) + 1
}

func min(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		return a
	}
	return b

}
