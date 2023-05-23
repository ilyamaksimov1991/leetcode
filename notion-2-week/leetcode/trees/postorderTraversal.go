package trees

// time O(n)
// memory O(n)
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)

	res = append(res, root.Val)
	return res
}

// time O(n)
// memory O(n)
func postorderTraversal2(root *TreeNode) []int {
	res := values(root)
	if res == nil {
		return nil
	}
	k, j := 0, len(res)-1
	for k < j {
		res[k], res[j] = res[j], res[k]
		k++
		j--

	}
	return res
}

func values(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, root.Val)
	res = append(res, values(root.Right)...)
	res = append(res, values(root.Left)...)
	return res
}
