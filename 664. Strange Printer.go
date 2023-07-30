package main

var dp map[[2]int]int

func strangePrinter(s string) int {
	dp = make(map[[2]int]int, len(s)*len(s))
	return minNumberOfTurns(s[len(s)-1], 0, len(s)-1, s) + 1

}

func minNumberOfTurns(c byte, i, j int, s string) int {
	if _, ok := dp[[2]int{i, j}]; ok {
		return dp[[2]int{i, j}]
	}

	k := -1

	for m := i; m <= j; m++ {
		if c != s[m] {
			k = m
			break
		}
	}

	if k == -1 {
		dp[[2]int{i, j}] = 0
		return 0
	}

	ans := minNumberOfTurns(s[k], k, k, s) + minNumberOfTurns(c, k+1, j, s)
	if c != s[k] {
		ans += 1
	}

	for m := k + 1; m < j; m++ {
		t := minNumberOfTurns(s[m], k, m, s) + minNumberOfTurns(c, m+1, j, s)
		if c != s[m] {
			t += 1
		}
		ans = min664(t, ans)
	}
	dp[[2]int{i, j}] = ans
	return ans
}

func min664(a, b int) int {
	if a < b {
		return a
	}
	return b
}
