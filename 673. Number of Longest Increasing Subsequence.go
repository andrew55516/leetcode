package main

func findNumberOfLIS(nums []int) int {
	ans := 1
	max := 1
	dp := make([][2]int, len(nums))
	dp[len(nums)-1] = [2]int{1, 1}
	for i := len(nums) - 2; i >= 0; i-- {
		dp[i] = [2]int{1, 1}
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				if dp[i][0] < dp[j][0]+1 {
					dp[i] = [2]int{dp[j][0] + 1, dp[j][1]}
				} else if dp[i][0] == dp[j][0]+1 {
					dp[i][1] += dp[j][1]
				}
			}
		}
		if dp[i][0] > max {
			max = dp[i][0]
			ans = dp[i][1]
		} else if dp[i][0] == max {
			ans += dp[i][1]
		}

	}
	return ans
}
