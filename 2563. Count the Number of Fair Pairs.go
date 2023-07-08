package main

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	ans := 0
	j := len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		left := i + 1
		right := j
		for left < right {
			mid := left + (right-left)/2
			if nums[mid]+nums[i] >= lower {
				right = mid
			} else {
				left = mid + 1
			}
		}
		min := left
		left = i
		right = j
		for left < right {
			mid := right - (right-left)/2
			if nums[mid]+nums[i] <= upper {
				left = mid
			} else {
				right = mid - 1
			}
		}
		max := right
		if nums[i]+nums[min] < lower || nums[i]+nums[max] > upper {
			continue
		}
		ans += max - min + 1
	}
	return int64(ans)
}
