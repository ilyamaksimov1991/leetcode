package list

// time O(n)
// memory O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		}

		cur = cur.Next
	}

	return head
}
