package main

import "fmt"

type test []int

func main() {
	a := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(medianSlidingWindow(a, 3))
}
