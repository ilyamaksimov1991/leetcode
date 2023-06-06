package list

/// TODO обсудить с Владимиром
func deleteNode(node *ListNode) {
	*node = *node.Next
}
