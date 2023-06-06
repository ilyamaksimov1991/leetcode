package list

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// разделить на 2 списка
	var prev, fast, slow *ListNode = nil, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	// объединить 2 списка
	return merge(sortList(head), sortList(slow))
}

func merge(node1 *ListNode, node2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			cur.Next = node1
			node1 = node1.Next
		} else {
			cur.Next = node2
			node2 = node2.Next
		}

		cur = cur.Next
	}

	if node1 != nil {
		cur.Next = node1
	}
	if node2 != nil {
		cur.Next = node2
	}

	return dummy.Next
}
