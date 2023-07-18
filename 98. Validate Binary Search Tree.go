package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	_, _, ok := check98(root)
	return ok
}

func check98(root *TreeNode) (int, int, bool) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val, true
	}

	if root.Left != nil && root.Right != nil {
		minLeft, maxLeft, okLeft := check98(root.Left)
		minRight, maxRight, okRight := check98(root.Right)

		return min98(minLeft, minRight, root.Val),
			max98(maxLeft, maxRight, root.Val),
			okLeft && okRight && root.Val > maxLeft && root.Val < minRight
	}

	if root.Left != nil {
		minLeft, maxLeft, okLeft := check98(root.Left)
		return min98(minLeft, root.Val), max98(maxLeft, root.Val), okLeft && root.Val > maxLeft
	}

	minRight, maxRight, okRight := check98(root.Right)
	return min98(minRight, root.Val), max98(maxRight, root.Val), okRight && root.Val < minRight

}

func min98(in ...int) int {
	min := in[0]
	for i := 1; i < len(in); i++ {
		if in[i] < min {
			min = in[i]
		}
	}
	return min
}

func max98(in ...int) int {
	max := in[0]
	for i := 1; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		}
	}
	return max
}
