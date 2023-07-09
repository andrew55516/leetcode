package main

func largestVariance(s string) int {
	alph := make([]byte, 0)
	counter := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := counter[s[i]]; !ok {
			alph = append(alph, s[i])
		}
		counter[s[i]]++
	}
	ans := 0
	for i := 0; i < len(alph); i++ {
		for j := i + 1; j < len(alph); j++ {
			m := max2272(kadane2272(alph[i], alph[j], counter[alph[j]], s),
				kadane2272(alph[j], alph[i], counter[alph[i]], s))
			ans = max2272(ans, m)
		}
	}
	return ans
}

func kadane2272(a byte, b byte, bCount int, s string) int {
	localMaxDiff := 0
	GlobalMaxDiff := 0
	currCountA := 0
	currCountB := 0
	for i := 0; i < len(s); i++ {
		if s[i] == a {
			localMaxDiff++
			currCountA++
		}
		if s[i] == b {
			localMaxDiff--
			currCountB++
			bCount--
		}
		if localMaxDiff < 0 && bCount > 0 {
			localMaxDiff = 0
			currCountA = 0
			currCountB = 0
		}
		if localMaxDiff > GlobalMaxDiff && currCountA > 0 && currCountB > 0 {
			GlobalMaxDiff = localMaxDiff
		}
	}
	return GlobalMaxDiff
}

func max2272(a, b int) int {
	if a > b {
		return a
	}
	return b
}
