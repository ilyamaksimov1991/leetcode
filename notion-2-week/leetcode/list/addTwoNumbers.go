package list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	current := res
	i := 0
	plus := 0
	for (l1 != nil && l2 != nil) || plus != 0 {
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

		summ := val1 + val2 + plus
		println(summ)
		if val1+val2 >= 10 {
			plus = 1
			summ = summ - 10
		} else {
			plus = 0
		}
		//println(summ)

		current.Next = &ListNode{
			Val: summ,
		}
		current = current.Next
		i++
		println(i)
	}

	return res.Next
}
