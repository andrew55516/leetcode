package main

import "math"

func repairCars(ranks []int, cars int) int64 {
	maxRank := 0
	for _, r := range ranks {
		if r > maxRank {
			maxRank = r
		}
	}

	left := 0
	right := maxRank * cars * cars
	for left < right {
		mid := left + (right-left)/2
		if check2594(ranks, cars, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return int64(left)
}

func check2594(ranks []int, cars int, t int) bool {
	for _, v := range ranks {
		amount := int(math.Sqrt(float64(t / v)))
		cars -= amount
	}
	if cars <= 0 {
		return true
	}
	return false
}
