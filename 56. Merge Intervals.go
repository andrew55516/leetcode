package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ans[len(ans)-1][1] ||
			ans[len(ans)-1][0] <= intervals[i][1] && intervals[i][1] <= ans[len(ans)-1][1] {

			ans[len(ans)-1][0] = min56(ans[len(ans)-1][0], intervals[i][0])
			ans[len(ans)-1][1] = max56(ans[len(ans)-1][1], intervals[i][1])

		} else {
			ans = append(ans, intervals[i])
		}
	}

	return ans
}

func min56(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max56(a, b int) int {
	if a > b {
		return a
	}
	return b

}
