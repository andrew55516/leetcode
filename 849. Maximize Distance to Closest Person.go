package main

func maxDistToClosest(seats []int) int {
	prev := -1
	ans := 0
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if prev == -1 {
				ans = i
			} else {
				ans = max849(ans, (i-prev)/2)
			}
			prev = i
		}
	}
	ans = max849(ans, len(seats)-1-prev)
	return ans
}

func max849(a, b int) int {
	if a > b {
		return a
	}
	return b
}
