package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	t := make([]*TreeNode, 1)
	t[0] = root
	k := 0
	for {
		if k%2 == 1 {
			for i := 0; i < len(t)/2; i++ {
				t[i].Val, t[len(t)-i-1].Val = t[len(t)-i-1].Val, t[i].Val
			}
		}
		if t[0].Left == nil {
			break
		}
		next := make([]*TreeNode, 0)
		for _, node := range t {
			next = append(next, node.Left, node.Right)
		}
		t = next
		k++
	}
	return root
}
