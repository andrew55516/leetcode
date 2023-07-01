package main

import "container/heap"

type intHeap1962 []int

func (h intHeap1962) Len() int {
	return len(h)
}

func (h intHeap1962) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap1962) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap1962) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap1962) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minStoneSum(piles []int, k int) int {
	h := &intHeap1962{}

	for _, p := range piles {
		h.Push(p)
	}

	heap.Init(h)

	for i := 0; i < k; i++ {
		pile := heap.Pop(h)
		newPile := pile.(int) - pile.(int)/2
		heap.Push(h, newPile)
	}

	sum := 0
	for _, v := range *h {
		sum += v
	}

	return sum
}
