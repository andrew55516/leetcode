package main

func splitArray(nums []int, k int) int {
	left := nums[0]
	right := 0
	for _, v := range nums {
		right += v
		if v > left {
			left = v
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if check410(nums, mid, k) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func check410(nums []int, min, k int) bool {
	sum := 0
	count := 1
	for _, v := range nums {
		if sum+v > min {
			count++
			sum = 0
		}
		sum += v
	}
	return count <= k
}
