package main

import "sort"

func minCost(basket1 []int, basket2 []int) int64 {
	basket1Map := make(map[int]int, 0)
	basket2Map := make(map[int]int, 0)
	basketMap := make(map[int]int, 0)
	min := basket1[0]
	toSwap := make([]int, 0)
	for _, v := range basket1 {
		basketMap[v]++
		basket1Map[v]++
		if v < min {
			min = v
		}
	}
	for _, v := range basket2 {
		basketMap[v]++
		basket2Map[v]++
		if v < min {
			min = v
		}
	}
	for k, v := range basketMap {
		if v%2 != 0 {
			return -1
		}
		basketMap[k] = v / 2
		if basket1Map[k] != basketMap[k] {
			swaps := abs2561(basket1Map[k] - basketMap[k])
			for i := 0; i < swaps; i++ {
				toSwap = append(toSwap, k)
			}
		}
	}

	sort.Ints(toSwap)

	ans := 0

	for i := 0; i < len(toSwap)/2; i++ {
		cost := min2561(min*2, toSwap[i])
		ans += cost
	}

	return int64(ans)
}

func min2561(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs2561(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
