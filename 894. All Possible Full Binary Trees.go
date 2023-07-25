package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}

	dp := make([][]*TreeNode, n+1)
	dp[1] = append(dp[1], &TreeNode{Val: 0})
	for i := 3; i <= n; i += 2 {
		for l := 1; l < i-1; l += 2 {
			r := i - l - 1
			for _, left := range dp[l] {
				for _, right := range dp[r] {
					dp[i] = append(dp[i], &TreeNode{Val: 0, Left: left, Right: right})
				}
			}
		}
	}

	return dp[n]
}
