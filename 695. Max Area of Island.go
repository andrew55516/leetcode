package main

func maxAreaOfIsland(grid [][]int) int {
	seen := make(map[[2]int]struct{}, 0)
	ans := 0
	
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if _, ok := seen[[2]int{i, j}]; !ok && grid[i][j] == 1 {
				seen[[2]int{i, j}] = struct{}{}
				ans = max695(ans, dfs695(&grid, &seen, i, j))
			}
		}
	}
	return ans

}

func dfs695(grid *[][]int, seen *map[[2]int]struct{}, i, j int) int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := 1
	for _, dir := range directions {
		x, y := i+dir[0], j+dir[1]
		if valid(x, y, grid) {
			if _, ok := (*seen)[[2]int{x, y}]; !ok {
				(*seen)[[2]int{x, y}] = struct{}{}
				ans += dfs695(grid, seen, x, y)
			}
		}
	}
	return ans
}

func valid(i, j int, grid *[][]int) bool {
	return i >= 0 && i < len(*grid) && j >= 0 && j < len((*grid)[0]) && (*grid)[i][j] == 1
}

func max695(a, b int) int {
	if a > b {
		return a
	}
	return b

}
