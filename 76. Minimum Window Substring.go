package main

func minWindow(s string, t string) string {
	shouldBeMatched := make(map[byte]int, 0)
	for i := range t {
		shouldBeMatched[t[i]]++
	}

	ans := [2]int{0, len(s)}
	accepted := 0
	i := 0
	j := 0
	for j < len(s) {
		if _, ok := shouldBeMatched[s[j]]; ok {
			shouldBeMatched[s[j]]--
			if shouldBeMatched[s[j]] == 0 {
				accepted++
			}
		}

		for i <= j {
			if _, ok := shouldBeMatched[s[i]]; ok && shouldBeMatched[s[i]] >= 0 {
				break
			}
			if _, ok := shouldBeMatched[s[i]]; ok {
				shouldBeMatched[s[i]]++
			}
			i++
		}
		if accepted == len(shouldBeMatched) {
			if j-i < ans[1]-ans[0] {
				ans = [2]int{i, j}
			}
		}
		j++
	}

	if ans[1] == len(s) {
		return ""
	}
	return s[ans[0] : ans[1]+1]
}
