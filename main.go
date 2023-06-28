package main

func main() {
	pairSum(nil)

}

func sortedSquares(nums []int) []int {
	a1 := make([]int, 0)
	a2 := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			a1 = append(a1, nums[i]*nums[i])
		} else {
			a2 = append(a2, nums[i]*nums[i])
		}
	}

	i := 0
	j := len(a2) - 1
	k := 0
	for i < len(a1) && j >= 0 {
		if a1[i] < a2[j] {
			nums[k] = a1[i]
			i++
		} else {
			nums[k] = a2[j]
			j--
		}
		k++
	}

	for i < len(a1) {
		nums[k] = a1[i]
		i++
		k++
	}

	for j >= 0 {
		nums[k] = a2[j]
		j--
		k++
	}
	return nums

}
