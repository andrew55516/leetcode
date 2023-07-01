package main

import "container/heap"

type intHeap1167 []int

func (h intHeap1167) Len() int {
	return len(h)
}

func (h intHeap1167) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap1167) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap1167) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap1167) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func connectSticks(sticks []int) int {
	h := &intHeap1167{}

	for _, v := range sticks {
		h.Push(v)
	}
	heap.Init(h)
	sum := 0

	for h.Len() > 1 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		sum += x + y
		heap.Push(h, x+y)
	}

	return sum
}
