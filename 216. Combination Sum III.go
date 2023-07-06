package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	for i := 1; i <= 9; i++ {
		backtracking216(n, k, i, i+1, fmt.Sprintf("%d", i), &ans)
	}
	return ans
}

func backtracking216(n, k, sum, next int, path string, combinations *[][]int) {
	if sum == n && len(path) == k {
		temp := make([]int, len(path))
		for i := 0; i < len(path); i++ {
			temp[i] = int(path[i] - '0')
		}
		*combinations = append(*combinations, temp)
		return
	}

	if sum > n || len(path) == k {
		return
	}

	for i := next; i <= 9; i++ {
		backtracking216(n, k, sum+i, i+1, fmt.Sprintf("%s%d", path, i), combinations)
	}
}
