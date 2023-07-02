package main

func smallestDivisor(nums []int, threshold int) int {
	left := 1
	right := 1
	for _, v := range nums {
		if v > right {
			right = v
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if check1283(nums, mid, threshold) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check1283(nums []int, divisor, threshold int) bool {
	sum := 0
	for _, v := range nums {
		sum += v / divisor
		if v%divisor != 0 {
			sum += 1
		}
	}
	return sum <= threshold
}
