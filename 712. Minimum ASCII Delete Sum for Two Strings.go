package main

var cache [][]int

func minimumDeleteSum(s1 string, s2 string) int {
	isEqual := make([][]bool, len(s1))
	cache = make([][]int, len(s1))
	for i := range s1 {
		cache[i] = make([]int, len(s2))
		isEqual[i] = make([]bool, len(s2))
		for j := range s2 {
			cache[i][j] = -1
			if s1[i] == s2[j] {
				isEqual[i][j] = true
			}
		}
	}

	sumOfAll := 0
	for i := 0; i < len(s1); i++ {
		sumOfAll += int(s1[i])
	}
	for i := 0; i < len(s2); i++ {
		sumOfAll += int(s2[i])
	}

	return sumOfAll - maxCommonSS(s1, s2, 0, 0, isEqual)
}

func maxCommonSS(s1 string, s2 string, ind1, ind2 int, m [][]bool) int {
	if ind1 >= len(s1) || ind2 >= len(s2) {
		return 0
	}

	if cache[ind1][ind2] != -1 {
		return cache[ind1][ind2]
	}

	ans := 0
	for i := ind2; i < len(s2); i++ {
		if m[ind1][i] {
			ans = max712(ans, int(s2[i])+int(s1[ind1])+maxCommonSS(s1, s2, ind1+1, i+1, m))
		}
	}
	ans = max712(ans, maxCommonSS(s1, s2, ind1+1, ind2, m))

	cache[ind1][ind2] = ans
	return ans
}

func max712(a, b int) int {
	if a > b {
		return a
	}
	return b
}
