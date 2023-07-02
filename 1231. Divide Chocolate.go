package main

func maximizeSweetness(sweetness []int, k int) int {
	left := sweetness[0]
	right := 0
	for _, v := range sweetness {
		right += v
		if v < left {
			left = v
		}
	}

	right = right / (k + 1)

	for left < right {
		mid := right - (right-left)/2
		if check1231(sweetness, mid, k) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left

}

func check1231(sweetness []int, min, k int) bool {
	sum := 0
	count := 0
	for _, v := range sweetness {
		sum += v
		if sum >= min {
			count++
			sum = 0
		}
	}
	return count >= k+1
}
