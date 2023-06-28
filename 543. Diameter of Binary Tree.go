package main

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diameter := 0

	dfs1(root, &diameter)

	return diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func dfs1(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	left := dfs1(root.Left, diameter)
	right := dfs1(root.Right, diameter)

	*diameter = max(*diameter, left+right)

	return max(left, right) + 1
}
