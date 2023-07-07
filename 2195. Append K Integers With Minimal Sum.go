package main

import "sort"

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	ans := 0
	if nums[0] > 1 {
		if nums[0]-1 >= k {
			ans += (1 + k) * k / 2
			k = 0
		} else {
			ans += (1 + nums[0] - 1) * (nums[0] - 1) / 2
			k -= nums[0] - 1
		}
	}
	for i := 1; i < len(nums) && k > 0; i++ {
		if nums[i]-nums[i-1]-1 >= k {
			ans += (nums[i-1] + 1 + nums[i-1] + k) * k / 2
			k = 0
		} else if nums[i]-nums[i-1] > 0 {
			ans += (nums[i-1] + 1 + nums[i] - 1) * (nums[i] - nums[i-1] - 1) / 2
			k -= nums[i] - nums[i-1] - 1
		}
	}

	if k > 0 {
		ans += (nums[len(nums)-1] + 1 + nums[len(nums)-1] + k) * k / 2
	}
	return int64(ans)
}
