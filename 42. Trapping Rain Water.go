package main

func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		if i == 0 {
			left[i] = height[i]
		} else {
			left[i] = max42(left[i-1], height[i])
		}
	}

	for i := len(height) - 1; i >= 0; i-- {
		if i == len(height)-1 {
			right[i] = height[i]
		} else {
			right[i] = max42(right[i+1], height[i])
		}
	}

	ans := 0

	for i := 0; i < len(height); i++ {
		ans += min42(left[i], right[i]) - height[i]
	}

	return ans
}

func min42(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max42(a, b int) int {
	if a > b {
		return a
	}
	return b

}
