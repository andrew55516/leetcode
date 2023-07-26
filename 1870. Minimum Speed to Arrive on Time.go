package main

import "math"

func minSpeedOnTime(dist []int, hour float64) int {
	h := hour - float64(len(dist)-1)

	if h <= 0 {
		return -1
	}

	left := 1
	right := int(float64(dist[len(dist)-1])/h) + 1
	for i := 0; i < len(dist); i++ {
		if dist[i] > right {
			right = dist[i]
		}
	}

	for left < right {
		mid := (left + right) / 2
		if check1870(dist, hour, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check1870(dist []int, hour float64, speed int) bool {
	t := 0
	for i := 0; i < len(dist)-1; i++ {
		t += dist[i] / speed
		if dist[i]%speed > 0 {
			t++
		}
	}
	k := int(math.Round(hour * 100))

	h := (k - t*100) * speed
	return dist[len(dist)-1]*100 <= h
}
