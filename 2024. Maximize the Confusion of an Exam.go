package main

func maxConsecutiveAnswers(answerKey string, k int) int {
	ans := 0
	left := 0
	right := 0
	t := 0
	f := 0
	for right < len(answerKey) {
		if answerKey[right] == 'T' {
			t++
		} else {
			f++
		}

		if t > k && f > k {
			if answerKey[left] == 'T' {
				t--
			} else {
				f--
			}
			left++
		}
		ans = max2024(ans, right-left+1)
		right++
	}
	return ans
}

func max2024(a, b int) int {
	if a > b {
		return a
	}
	return b
}
