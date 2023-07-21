package main

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs200(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs200(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs200(grid, i+1, j)
	dfs200(grid, i-1, j)
	dfs200(grid, i, j+1)
	dfs200(grid, i, j-1)
}
