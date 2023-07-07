package main

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	right := 0
	ans := len(nums) + 1
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			ans = min209(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func min209(a, b int) int {
	if a < b {
		return a
	}
	return b
}
