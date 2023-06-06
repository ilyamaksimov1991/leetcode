package list

func mergeKLists(lists []*ListNode) *ListNode {
	list1 := &ListNode{}
	cur1 := list1

	for i := 0; i < len(lists); i++ {
		cur1.Next = lists[i]
		for cur1.Next != nil {
			cur1 = cur1.Next
		}
	}

	return sortList(list1.Next)
}
