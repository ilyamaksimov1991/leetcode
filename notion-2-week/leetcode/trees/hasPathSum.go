package trees

// time O(n)
// memory O(1)
func hasPathSum(root *TreeNode, targetSum int) bool {
	switch {
	case root == nil:
		return false
	case targetSum == root.Val && root.Right == nil && root.Left == nil:
		return true
	case hasPathSum(root.Left, targetSum-root.Val):
		return true
	case hasPathSum(root.Right, targetSum-root.Val):
		return true
	}
	return false
}

// time O(n)
// memory O(1)
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	l := hasPathSum(root.Left, targetSum-root.Val)
	if l {
		return l
	}

	r := hasPathSum(root.Right, targetSum-root.Val)
	if r {
		return r
	}

	return false
}
func hasPathSum3(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum = targetSum - root.Val

	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	l := hasPathSum(root.Left, targetSum)
	if l {
		return l
	}

	r := hasPathSum(root.Right, targetSum)
	if r {
		return r
	}

	return false
}
