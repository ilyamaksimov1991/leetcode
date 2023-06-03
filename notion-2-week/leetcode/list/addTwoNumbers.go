package list

// time O(n)
// memory O(n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	dummy := res
	curry := 0
	for l1 != nil || l2 != nil || curry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + curry
		curry = sum / 10

		dummy.Next = &ListNode{
			Val: sum % 10,
		}
		dummy = dummy.Next
	}

	return res.Next
}
