package main

import "container/heap"

type maxHeap239 []*item239

type item239 struct {
	val   int
	index int
}

func (h maxHeap239) Len() int {
	return len(h)
}

func (h maxHeap239) Less(i, j int) bool {
	return h[i].val > h[j].val
}

func (h maxHeap239) Swap(i, j int) {
	h[i].index = j
	h[j].index = i
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap239) Push(x any) {
	item := x.(*item239)
	item.index = len(*h)
	*h = append(*h, item)
}

func (h *maxHeap239) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	(*h)[n-1] = nil
	*h = (*h)[:n-1]
	return x
}

func (h *maxHeap239) Peek() *item239 {
	return (*h)[0]
}

func (h *maxHeap239) update(item *item239, val int) {
	item.val = val
	heap.Fix(h, item.index)
}

func maxSlidingWindow(nums []int, k int) []int {
	h := new(maxHeap239)
	heap.Init(h)

	ans := make([]int, 0)
	items := make([]*item239, 0)
	for i := 0; i < k; i++ {
		item := &item239{val: nums[i]}
		items = append(items, item)
		heap.Push(h, item)
	}

	ans = append(ans, h.Peek().val)

	for i := k; i < len(nums); i++ {
		item := items[0]
		h.update(item, nums[i])
		items = items[1:]
		items = append(items, item)
		ans = append(ans, h.Peek().val)
	}

	return ans
}
