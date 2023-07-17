package main

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int, 0)
	ans := [2]int{0, 0}
	last := -1
	for i := 0; i < len(s); i++ {
		if _, ok := seen[s[i]]; ok {
			last = max3(last, seen[s[i]])
		}
		if ans[1]-ans[0] < i-last {
			ans = [2]int{last, i}
		}

		seen[s[i]] = i
	}
	return ans[1] - ans[0]
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
