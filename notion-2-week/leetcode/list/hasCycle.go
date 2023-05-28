package list

// time O(n)
// memory O(n)
func hasCycle(head *ListNode) bool {
	addr := make(map[*ListNode]bool)
	i := 0
	for head != nil {
		if _, ok := addr[head]; ok {
			return true
		} else {
			addr[head] = true
		}
		i++
		head = head.Next
	}
	return false
}

// time O(n)
// memory O(1)
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow == fast
}
