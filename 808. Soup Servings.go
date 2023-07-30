package main

func soupServings(n int) float64 {
	k := n/25 + 1
	if n%25 != 0 {
		k++
	}

	dp := make(map[int]map[int]float64, 0)
	dp[0] = make(map[int]float64, 0)
	dp[0][0] = 0.5

	serves := [][2]int{
		{4, 0},
		{3, 1},
		{2, 2},
		{1, 3},
	}

	for i := 1; i < k; i++ {
		dp[i] = make(map[int]float64, 0)
		dp[0][i] = 1
		dp[i][0] = 0
		for j := 1; j <= i; j++ {
			for _, serve := range serves {
				dp[i][j] += 0.25 * dp[max808(i-serve[0], 0)][max808(j-serve[1], 0)]
				if i != j {
					dp[j][i] += 0.25 * dp[max808(j-serve[0], 0)][max808(i-serve[1], 0)]
				}
			}
		}
		if dp[i][i] > 1-1e-5 {
			return 1
		}
	}

	return dp[k-1][k-1]
}

func max808(a, b int) int {
	if a > b {
		return a
	}
	return b
}
