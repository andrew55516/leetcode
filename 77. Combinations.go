package main

import (
	"strconv"
	"strings"
)

var combinations [][]int

func combine(n int, k int) [][]int {
	combinations = make([][]int, 0)
	comb(n, k, k, 1, "")
	return combinations
}

func comb(n, k, l, i int, s string) {
	if l == 0 {
		ans := make([]int, k)
		for j, v := range strings.Split(strings.TrimSpace(s), " ") {
			ans[j], _ = strconv.Atoi(v)
		}
		combinations = append(combinations, ans)
		return
	}

	if i > n {
		return
	}
	comb(n, k, l-1, i+1, s+" "+strconv.Itoa(i))
	comb(n, k, l, i+1, s)
}
