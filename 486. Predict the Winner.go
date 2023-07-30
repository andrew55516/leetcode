package main

func PredictTheWinner(nums []int) bool {
	//cache := make(map[int]int, 0)

	m1, m2 := dp486(nums, 0)

	if m1 >= m2 {
		return true
	}
	return false
}

func dp486(nums []int, k int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}

	m1, m2 := dp486(nums[1:], k+1)
	n1, n2 := dp486(nums[:len(nums)-1], k+1)

	if k%2 == 0 {
		m1 += nums[0]
		n1 += nums[len(nums)-1]
		if m1 >= n1 {
			return m1, m2
		} else {
			return n1, n2
		}
	} else {
		m2 += nums[0]
		n2 += nums[len(nums)-1]
		if m2 >= n2 {
			return m1, m2
		} else {
			return n1, n2
		}
	}
}
