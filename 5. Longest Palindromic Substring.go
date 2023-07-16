package main

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	ans := [2]int{0, 0}
	m := make([][]bool, len(s))
	for i := range m {
		m[i] = make([]bool, len(s))
		m[i][i] = true
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			ans = [2]int{i, i + 1}
			m[i][i+1] = true
		}
	}

	for d := 2; d < len(s); d++ {
		for i := 0; i < len(s)-d; i++ {
			if s[i] == s[i+d] && m[i+1][i+d-1] {
				m[i][i+d] = true
				ans = [2]int{i, i + d}
			}
		}
	}

	return s[ans[0] : ans[1]+1]
}
