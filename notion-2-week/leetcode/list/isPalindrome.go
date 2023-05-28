package list

// time O(n)
// memory O(1) ?
func isPalindrome(head *ListNode) bool {
	// находим середину
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// разворачиваем середину
	cur := slow
	var rev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = rev
		rev = cur
		cur = next
	}

	// проверяем с начала и с конца
	for rev != nil {
		if rev.Val != head.Val {
			return false
		}
		rev = rev.Next
		head = head.Next
	}

	return true
}
