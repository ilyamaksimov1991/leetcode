package trees

// TODO доделать
func isSymmetric(root *TreeNode) bool {
	if root.Left != nil && root.Right != nil {
		return isSymmetric(root.Left) == isSymmetric(root.Right)
	}

	return false
}
