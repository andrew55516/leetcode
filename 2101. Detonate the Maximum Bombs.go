package main

var bombsMap map[int][]int

func maximumDetonation(bombs [][]int) int {
	ans := 0
	bombsMap = make(map[int][]int, len(bombs))

	for i := 0; i < len(bombs); i++ {
		for j := i + 1; j < len(bombs); j++ {
			bomb1 := bombs[i]
			bomb2 := bombs[j]
			dx := bomb1[0] - bomb2[0]
			dy := bomb1[1] - bomb2[1]
			r := dx*dx + dy*dy
			if r <= bomb1[2]*bomb1[2] {
				bombsMap[i] = append(bombsMap[i], j)
			}
			if r <= bomb2[2]*bomb2[2] {
				bombsMap[j] = append(bombsMap[j], i)
			}
		}
	}

	//seen = make([]bool, len(bombs))
	for i := 0; i < len(bombs); i++ {
		//if !seen[i] {
		seen := make([]bool, len(bombs))
		seen[i] = true
		ans = max2101(ans, dfs2101(i, seen))
		//}
	}

	return ans
}

func dfs2101(start int, seen []bool) int {
	ans := 1

	for _, bomb := range bombsMap[start] {
		if !seen[bomb] {
			seen[bomb] = true
			ans += dfs2101(bomb, seen)
		}
	}

	return ans
}

func max2101(a, b int) int {
	if a > b {
		return a
	}
	return b
}
