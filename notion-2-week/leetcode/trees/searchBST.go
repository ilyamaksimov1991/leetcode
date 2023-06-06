package trees

// time O(n)
// memory O(1)
func searchBST(root *TreeNode, val int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val == val:
		return root
	case root.Val > val:
		return searchBST(root.Left, val)
	case root.Val < val:
		return searchBST(root.Right, val)
	}

	return nil
}

// time O(n)
// memory O(1)
func searchBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	l := searchBST(root.Left, val)
	if l != nil {
		return l
	}
	r := searchBST(root.Right, val)
	if r != nil {
		return r
	}

	return nil
}
