package main

import "container/heap"

type minHeap215 []int

func (h minHeap215) Len() int {
	return len(h)
}

func (h minHeap215) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap215) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap215) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap215) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	minHeap := &minHeap215{}
	heap.Init(minHeap)
	for i := 0; i < k; i++ {
		heap.Push(minHeap, nums[i])
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > (*minHeap)[0] {
			heap.Pop(minHeap)
			heap.Push(minHeap, nums[i])
		}
	}
	return (*minHeap)[0]
}
