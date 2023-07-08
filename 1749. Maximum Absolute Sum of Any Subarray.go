package main

func maxAbsoluteSum(nums []int) int {
	maxPrefix := 0
	minPrefix := 0
	curr := 0
	for i := 0; i < len(nums); i++ {
		curr += nums[i]
		if curr > maxPrefix {
			maxPrefix = curr
		}
		if curr < minPrefix {
			minPrefix = curr
		}
	}

	return maxPrefix - minPrefix
}
