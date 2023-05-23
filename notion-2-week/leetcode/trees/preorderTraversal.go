package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time O(n)
// memory O(n)
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, root.Val)

	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)

	return res
}
