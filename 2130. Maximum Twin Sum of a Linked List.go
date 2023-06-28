package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow := head
	fast := head
	k := 0
	maxSum := 0
	var preMiddle *ListNode

	for fast != nil && fast.Next != nil {
		k += 1
		preMiddle = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	curr := slow

	var prev *ListNode
	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}

	if preMiddle != nil {
		preMiddle.Next = prev
	}

	slow = head
	fast = head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		if slow.Val+fast.Val > maxSum {
			maxSum = slow.Val + fast.Val
		}
		fast = fast.Next
		slow = slow.Next
	}

	return maxSum
}
