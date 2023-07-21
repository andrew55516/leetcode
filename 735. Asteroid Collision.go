package main

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, v := range asteroids {
		if len(stack) == 0 || v > 0 {
			stack = append(stack, v)
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > 0 && v < 0 {
			m := stack[len(stack)-1]
			switch {
			case m == -v:
				v = 0
				stack = stack[:len(stack)-1]
			case m > -v:
				v = 0
			case m < -v:
				stack = stack[:len(stack)-1]
			}
		}

		if v < 0 {
			stack = append(stack, v)
		}
	}
	return stack
}
