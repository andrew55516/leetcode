package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	ans := 0
	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if last[1] > intervals[i][0] {
			ans++
		} else {
			last = intervals[i]
		}
	}

	return ans
}
