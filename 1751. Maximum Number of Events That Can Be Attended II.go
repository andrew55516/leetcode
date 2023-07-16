package main

import "sort"

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	cache := make([][]int, len(events))
	for i := range cache {
		cache[i] = make([]int, k+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	next := make([]int, len(events))
	for i := range next {
		next[i] = bs1751(events, i)
	}

	return dp1751(events, cache, next, 0, k)
}

func dp1751(events [][]int, cache [][]int, next []int, i, k int) int {
	if k == 0 || i == len(events) {
		return 0
	}

	if cache[i][k] != -1 {
		return cache[i][k]
	}

	cache[i][k] = max1751(events[i][2]+dp1751(events, cache, next, next[i], k-1),
		dp1751(events, cache, next, i+1, k))

	return cache[i][k]
}

func max1751(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bs1751(events [][]int, i int) int {
	left, right := i+1, len(events)
	for left < right {
		mid := left + (right-left)/2
		if events[mid][0] > events[i][1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left >= len(events) || events[left][0] <= events[i][1] {
		return len(events)
	}
	return left
}
