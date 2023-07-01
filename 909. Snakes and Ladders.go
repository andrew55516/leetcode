package main

func snakesAndLadders(board [][]int) int {
	n := len(board)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		board[i], board[j] = board[j], board[i]
		for _, m := range []int{i, j} {
			if m%2 == 1 {
				for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
					board[m][j], board[m][k] = board[m][k], board[m][j]
				}
			}
		}
	}

	if n%2 == 1 {
		m := n / 2
		if m%2 == 1 {
			for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
				board[m][j], board[m][k] = board[m][k], board[m][j]
			}
		}
	}

	seen := make(map[[2]int]struct{}, 0)
	seen[[2]int{0, 0}] = struct{}{}

	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0, 0})

	for len(queue) > 0 {
		x, y, steps := queue[0][0], queue[0][1], queue[0][2]
		if board[x][y] != -1 {
			cel := board[x][y] - 1
			x, y = cel/n, cel%n
			//seen[[2]int{x, y}] = struct{}{}
		}
		cel := x*n + (y + 1)
		if cel == n*n {
			return steps
		}
		queue = queue[1:]

		for i := 1; i <= 6; i++ {
			nextCel := cel + i
			if nextCel <= n*n {
				x, y := (nextCel-1)/n, (nextCel-1)%n
				if _, ok := seen[[2]int{x, y}]; !ok {
					queue = append(queue, []int{x, y, steps + 1})
					seen[[2]int{x, y}] = struct{}{}
				}
			}
		}
	}

	return -1
}
