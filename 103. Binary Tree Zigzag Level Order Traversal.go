package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	if root == nil {
		return ans
	}

	queue := make([]*TreeNode, 0)

	queue = append(queue, root)

	i := 0

	for len(queue) > 0 {
		k := len(queue)
		vals := make([]int, 0)

		for j := 0; j < k; j++ {
			vals = append(vals, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}

			queue = queue[1:]

		}

		if i%2 == 1 {
			for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
				vals[i], vals[j] = vals[j], vals[i]
			}
		}

		ans = append(ans, vals)
		i++
	}

	return ans
}
