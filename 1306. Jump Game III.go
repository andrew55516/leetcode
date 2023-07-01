package main

func canReach(arr []int, start int) bool {
	seen := make([]bool, len(arr))
	seen[start] = true

	stack := []int{start}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if arr[curr] == 0 {
			return true
		}

		for _, v := range []int{-1, 1} {
			next := curr + v*arr[curr]
			if next >= 0 && next < len(arr) && !seen[next] {
				stack = append(stack, next)
				seen[next] = true
			}
		}
	}

	return false
}
