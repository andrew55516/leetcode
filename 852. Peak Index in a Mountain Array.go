package main

func peakIndexInMountainArray(arr []int) int {
	return mergeMax(arr, 0, len(arr)-1)
}

func mergeMax(arr []int, a, b int) int {
	if a >= b {
		return a
	}

	half := (a + b) / 2

	m1 := mergeMax(arr, a, half)
	m2 := mergeMax(arr, half+1, b)
	if arr[m1] > arr[m2] {
		return m1
	}
	return m2
}
