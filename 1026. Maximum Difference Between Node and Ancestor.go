package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := root.Val
	min := root.Val

	return dfs(root, max, min)

}

func dfs(root *TreeNode, max, min int) int {
	if root == nil {
		return max - min
	}

	if root.Val > max {
		max = root.Val
	}

	if root.Val < min {
		min = root.Val
	}

	left := dfs(root.Left, max, min)
	right := dfs(root.Right, max, min)

	if left > right {
		return left
	}

	return right
}
