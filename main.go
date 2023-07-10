package main

import "fmt"

type test []int

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(NonEmptyLevels(height))
}
