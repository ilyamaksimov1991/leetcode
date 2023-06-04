package trees

import "math"

// https://www.youtube.com/watch?v=s6ATEkipzow
// time O(n)
// memory O(1)
func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
}
