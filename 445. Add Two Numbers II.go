package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	r1 := l1
	var prev *ListNode
	for r1 != nil {
		next := r1.Next
		r1.Next = prev
		prev = r1
		r1 = next
	}
	r1 = prev

	r2 := l2
	prev = nil
	for r2 != nil {
		next := r2.Next
		r2.Next = prev
		prev = r2
		r2 = next
	}
	r2 = prev

	root := &ListNode{Val: r1.Val + r2.Val}
	add := root.Val / 10
	root.Val %= 10

	r1 = r1.Next
	r2 = r2.Next

	for r1 != nil || r2 != nil {
		curr := &ListNode{Val: add}
		if r1 != nil {
			curr.Val += r1.Val
			r1 = r1.Next
		}
		if r2 != nil {
			curr.Val += r2.Val
			r2 = r2.Next
		}

		add = curr.Val / 10
		curr.Val %= 10

		curr.Next = root
		root = curr
	}

	if add != 0 {
		curr := &ListNode{Val: add, Next: root}
		root = curr
	}
	return root
}
