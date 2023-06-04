package trees

import (
	"fmt"
	"strconv"
)

// time O(n)
// memory O(n)
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	return recursive(root, "")
}

func recursive(node *TreeNode, str string) []string {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return []string{str + strconv.Itoa(node.Val)}
	}

	left := recursive(node.Left, str+fmt.Sprintf("%d->", node.Val))
	right := recursive(node.Right, str+fmt.Sprintf("%d->", node.Val))

	return append(left, right...)
}
