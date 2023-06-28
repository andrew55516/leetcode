package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	ansMap := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		ansMap[v] = -1
	}

	decrStack := make([]int, 0)
	for _, v := range nums2 {
		for len(decrStack) > 0 && decrStack[len(decrStack)-1] < v {
			if _, ok := ansMap[decrStack[len(decrStack)-1]]; ok {
				ansMap[decrStack[len(decrStack)-1]] = v
			}
			decrStack = decrStack[:len(decrStack)-1]
		}
		decrStack = append(decrStack, v)
	}

	for i, v := range nums1 {
		ans[i] = ansMap[v]
	}

	return ans
}
