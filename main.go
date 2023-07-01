package main

import "fmt"

func main() {
	maze := [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}
	entrance := []int{1, 2}
	fmt.Println(maze[1][2] == '.')
	fmt.Println(nearestExit(maze, entrance))

}
