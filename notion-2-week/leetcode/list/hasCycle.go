package list

func hasCycle(head *ListNode) bool {
	println(head)
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
