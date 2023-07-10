package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	currL1 := l1.Next
	currL2 := l2.Next
	root := &ListNode{Val: l1.Val + l2.Val}
	add := root.Val / 10
	root.Val %= 10
	curr := root

	for currL1 != nil || currL2 != nil {
		curr.Next = &ListNode{Val: add}
		curr = curr.Next
		if currL1 != nil {
			curr.Val += currL1.Val
			currL1 = currL1.Next
		}
		if currL2 != nil {
			curr.Val += currL2.Val
			currL2 = currL2.Next
		}
		add = curr.Val / 10
		curr.Val %= 10
	}

	if add != 0 {
		curr.Next = &ListNode{Val: add}
	}
	return root
}
