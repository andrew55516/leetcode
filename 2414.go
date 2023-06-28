package main

func longestContinuousSubstring(s string) int {
	var max, k int
	var last int32
	for _, c := range s {
		if max == 0 {
			max = 1
			k = 1
		} else {
			if c == last+1 {
				k += 1
			} else {
				if k > max {
					max = k
				}
				k = 1
			}
		}
		last = c
	}
	if k > max {
		max = k
	}
	return max
}
