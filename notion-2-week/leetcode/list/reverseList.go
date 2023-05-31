package list

// time O(n)
// memory O(1)
func reverseList(head *ListNode) *ListNode {
	var rev *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = rev
		rev = head
		head = tmp
	}

	return rev
}
