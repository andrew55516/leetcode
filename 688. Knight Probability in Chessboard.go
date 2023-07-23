package main

func knightProbability(n int, k int, row int, column int) float64 {
	seen := make(map[[3]int]float64)
	return dp688(n, k, row, column, &seen)
}

func dp688(n int, k int, row int, column int, seen *map[[3]int]float64) float64 {
	steps := [][2]int{
		{1, 2},
		{2, 1},
		{2, -1},
		{1, -2},
		{-1, -2},
		{-2, -1},
		{-2, 1},
		{-1, 2},
	}

	if k == 0 {
		return 1
	}
	if _, ok := (*seen)[[3]int{row, column, k}]; ok {
		return (*seen)[[3]int{row, column, k}]
	}

	ans := float64(0)
	for _, step := range steps {
		nextRow := row + step[0]
		nextColumn := column + step[1]
		if nextRow >= 0 && nextRow < n && nextColumn >= 0 && nextColumn < n {
			ans += 0.125 * dp688(n, k-1, nextRow, nextColumn, seen)
		}
	}
	(*seen)[[3]int{row, column, k}] = ans
	return ans
}
