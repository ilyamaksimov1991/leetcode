package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// time O(n)
// memory O(1)
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}

	return slow
}
