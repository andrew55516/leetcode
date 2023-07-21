package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		switch v {
		case "+":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
		case "-":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]-stack[len(stack)-1])
		case "*":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
		case "/":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]/stack[len(stack)-1])
		default:
			n, _ := strconv.Atoi(v)
			stack = append(stack, n)
		}
	}
	return stack[0]
}
