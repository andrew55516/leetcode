package main

type state111 struct {
	node  *TreeNode
	depth int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []state111{state111{root, 1}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.node.Left == nil && curr.node.Right == nil {
			return curr.depth
		}
		if curr.node.Left != nil {
			queue = append(queue, state111{curr.node.Left, curr.depth + 1})
		}
		if curr.node.Right != nil {
			queue = append(queue, state111{curr.node.Right, curr.depth + 1})
		}
	}
	return 0
}
