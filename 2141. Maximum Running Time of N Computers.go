package main

func maxRunTime(n int, batteries []int) int64 {
	left := 0
	sum := 0
	for i := 0; i < len(batteries); i++ {
		sum += batteries[i]
	}
	right := sum/n + 1

	for left < right {
		mid := right - (right-left)/2
		if check2141(n, mid, batteries) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return int64(right)

}

func check2141(n, target int, batteries []int) bool {
	sum := 0
	for i := 0; i < len(batteries); i++ {
		sum += min2141(batteries[i], target)
	}

	return sum/n >= target
}

func min2141(a, b int) int {
	if a < b {
		return a
	}
	return b
}
