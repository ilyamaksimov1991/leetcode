package trees

// time O(nLogn)
// memory O(1)
func isSymmetric(root *TreeNode) bool {
	return dsf(root.Left, root.Right)
}

func dsf(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && dsf(left.Left, right.Right) && dsf(left.Right, right.Left)
}
