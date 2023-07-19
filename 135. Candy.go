package main

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}

	candies := make([]int, len(ratings))
	candies[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max135(candies[i], candies[i+1]+1)
		}
	}

	ans := 0
	for i := 0; i < len(candies); i++ {
		ans += candies[i]
	}
	return ans
}

func max135(a, b int) int {
	if a > b {
		return a
	}
	return b
}
