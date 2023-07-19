package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := root.Val
	checkSum(root, &ans)
	return ans
}

func checkSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	leftSum := checkSum(root.Left, maxSum)
	rightSum := checkSum(root.Right, maxSum)
	m := max124(root.Val, root.Val+leftSum, root.Val+rightSum)
	*maxSum = max124(*maxSum, m, leftSum+rightSum+root.Val)
	return m
}

func max124(a ...int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}
