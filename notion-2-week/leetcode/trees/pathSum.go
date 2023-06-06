package trees

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return append(res, []int{root.Val})
		}
		return res
	}

	println(targetSum - root.Val)
	for _, path := range pathSum(root.Left, targetSum-root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}

	for _, path := range pathSum(root.Right, targetSum-root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}

	return res
}
