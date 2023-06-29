package main

import "fmt"

func main() {
	grid := [][]int{{1, 1, 0, 1, 1}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {1, 1, 0, 1, 1}}
	fmt.Println(maxAreaOfIsland(grid))

}
