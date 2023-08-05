package main

import (
	"strconv"
	"strings"
)

var permutations [][]int

func permute(nums []int) [][]int {
	used := 0
	for i := 0; i < len(nums); i++ {
		used = used | (1 << i)
	}

	permutations = make([][]int, 0)
	perm(nums, used, "")
	return permutations
}

func perm(nums []int, used int, p string) {
	if used == 0 {
		permutation := make([]int, len(nums))
		for i, v := range strings.Split(strings.TrimSpace(p), " ") {
			permutation[i], _ = strconv.Atoi(v)
		}
		permutations = append(permutations, permutation)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used&(1<<i) == 0 {
			continue
		}
		perm(nums, used-(1<<i), p+" "+strconv.Itoa(nums[i]))
	}
}
