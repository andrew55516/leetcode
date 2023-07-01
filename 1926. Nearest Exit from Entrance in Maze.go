package main

var m, n int

func nearestExit(maze [][]byte, entrance []int) int {
	queue := make([][]int, 0)
	queue = append(queue, append(entrance, 0))
	seen := make(map[[2]int]struct{}, 0)
	seen[[2]int{entrance[0], entrance[1]}] = struct{}{}
	m = len(maze)
	n = len(maze[0])

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		k := len(queue)
		for i := 0; i < k; i++ {
			row, col, steps := queue[0][0], queue[0][1], queue[0][2]
			if isEntrance(row, col) && (row != entrance[0] || col != entrance[1]) {
				return steps
			}
			queue = queue[1:]
			for _, d := range directions {
				nextRow, nextCol := row+d[0], col+d[1]
				if _, ok := seen[[2]int{nextRow, nextCol}]; !ok && valid1926(nextRow, nextCol) && maze[nextRow][nextCol] == '.' {
					seen[[2]int{nextRow, nextCol}] = struct{}{}
					queue = append(queue, []int{nextRow, nextCol, steps + 1})
				}
			}

		}
	}

	return -1
}

func isEntrance(row, col int) bool {
	return row == 0 || row == m-1 || col == 0 || col == n-1
}

func valid1926(row, col int) bool {
	return row >= 0 && row < m && col >= 0 && col < n
}
