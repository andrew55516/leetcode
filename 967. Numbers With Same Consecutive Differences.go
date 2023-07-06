package main

import "fmt"

func numsSameConsecDiff(n int, k int) []int {
	ans := make([]int, 0)
	for i := 1; i <= 9; i++ {
		backtraking967(n, k, i, &ans)
	}
	return ans
}

func backtraking967(n, k, num int, combinations *[]int) {
	if len(fmt.Sprintf("%d", num)) == n {
		*combinations = append(*combinations, num)
		return
	}

	right := num%10 + k
	if right < 10 {
		backtraking967(n, k, num*10+right, combinations)
	}

	left := num%10 - k
	if left >= 0 && left != right {
		backtraking967(n, k, num*10+left, combinations)
	}
}
