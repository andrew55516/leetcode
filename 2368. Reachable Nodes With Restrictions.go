package main

func reachableNodes(n int, edges [][]int, restricted []int) int {
	restrictedMap := make(map[int]struct{}, len(restricted))
	for _, v := range restricted {
		restrictedMap[v] = struct{}{}
	}

	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	seen := make([]bool, n)
	seen[0] = true
	return dfs2368(0, graph, seen, restrictedMap)
}

func dfs2368(node int, graph [][]int, seen []bool, restricted map[int]struct{}) int {
	if _, ok := restricted[node]; ok {
		return 0
	}

	ans := 1
	for _, v := range graph[node] {
		if !seen[v] {
			seen[v] = true
			ans += dfs2368(v, graph, seen, restricted)
		}
	}
	return ans
}
