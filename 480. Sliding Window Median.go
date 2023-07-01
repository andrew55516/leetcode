package main

import "container/heap"

type minHeap480 []int

func (h minHeap480) Len() int {
	return len(h)
}

func (h minHeap480) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap480) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap480) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap480) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap480 []int

func (h maxHeap480) Len() int {
	return len(h)
}

func (h maxHeap480) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap480) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap480) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap480) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func medianSlidingWindow(nums []int, k int) []float64 {
	min := &minHeap480{}
	max := &maxHeap480{}
	
	ans := make([]float64, 0, len(nums)-k+1)

	for i := 0; i < k; i++ {
		heap.Push(max, nums[i])
		heap.Push(min, heap.Pop(max))
		if min.Len() > max.Len() {
			heap.Push(max, heap.Pop(min))
		}
	}

	if min.Len() == max.Len() {
		ans = append(ans, float64((*min)[0]+(*max)[0])/2)
	} else {
		ans = append(ans, float64((*max)[0]))
	}

	for i := k; i < len(nums); i++ {
		x := nums[i-k]
		y := nums[i]
		ind := -1
		for j, v := range *max {
			if v == x {
				ind = j
				break
			}
		}

		if ind >= 0 {
			heap.Remove(max, ind)
		} else {
			for j, v := range *min {
				if v == x {
					heap.Remove(min, j)
					break
				}
			}
		}

		heap.Push(max, y)
		heap.Push(min, heap.Pop(max))
		if min.Len() > max.Len() {
			heap.Push(max, heap.Pop(min))
		}

		for max.Len() > min.Len()+1 {
			heap.Push(min, heap.Pop(max))
		}

		if min.Len() == max.Len() {
			ans = append(ans, float64((*min)[0]+(*max)[0])/2)
		} else {
			ans = append(ans, float64((*max)[0]))
		}
	}

	return ans
}
