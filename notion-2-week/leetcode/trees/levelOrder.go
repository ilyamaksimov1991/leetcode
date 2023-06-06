package trees

// time O(n)
// memory O(n)
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)
		resCurLevel := []int{}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			resCurLevel = append(resCurLevel, node.Val)
		}

		res = append(res, resCurLevel)
	}

	return res
}
