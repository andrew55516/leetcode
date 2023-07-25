package main

func myPow(x float64, n int) float64 {
	//if x == 1 {
	//	return 1
	//}
	//if x == -1 {
	//	return pow50(-1, n%2)
	//}
	if n >= 0 {
		return pow50(x, n)
	}
	return 1 / pow50(x, -n)
}

func pow50(x float64, n int) float64 {
	ans := float64(1)
	for n > 0 {
		curr := x
		k := 1
		for k*2 < n {
			curr *= curr
			k *= 2
		}
		ans *= curr
		n -= k
	}
	return ans
}
