package list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	var res *ListNode
	//list1,list2 = list1
	i := 0
	for list1 != nil && list2 != nil {
		val1 := 99
		val2 := 99
		if list1 != nil {
			val1 = list1.Val
		}
		if list2 != nil {
			val2 = list2.Val
		}

		var min int
		if val1 < val2 {
			min = val1
			list1 = list1.Next
		} else {
			min = val2
			list2 = list2.Next
		}
		println(val1, val2, min)
		node := &ListNode{
			Val: min,
		}
		if i == 0 {
			res = node
		} else {
			res.Next = node
			res = res
		}

		i++
	}

	return res
}
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}

	i := 0
	listOne, listTwo := list1, list2
	for listOne != nil || listTwo != nil {
		//if listOne == nil {
		//	res.Next = listTwo
		//	for res != nil {
		//		//println("we")
		//		println(res.Val)
		//		res = res.Next
		//	}
		//	return res
		//}
		//if listTwo == nil {
		//	for res != nil {
		//		//println("we")
		//		println(res.Val)
		//		res = res.Next
		//	}
		//	res.Next = listOne
		//	return res
		//}
		//println("1")

		var next *ListNode
		if listOne.Val > listTwo.Val {
			next = listOne
			listOne = listOne.Next
		} else {
			next = listTwo
			listTwo = listTwo.Next
		}
		if i == 0 {
			res = next
		}
		res.Next = next
		i++

		println(next.Val)

		//listOne, listTwo = listOne.Next, listTwo.Next
	}

	return res
}
