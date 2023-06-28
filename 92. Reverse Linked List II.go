package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var preLeft, prev, leftNode *ListNode

	curr := head

	for i := 1; i <= right; i++ {
		switch {

		case i == left-1:
			preLeft = curr

		case i == left:
			leftNode = curr
			prev = curr

		case i > left:
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		if i <= left {
			curr = curr.Next
		}
	}

	if preLeft != nil {
		preLeft.Next = prev
	} else {
		head = prev
	}

	leftNode.Next = curr

	return head
}
