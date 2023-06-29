package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	seen := make([]bool, n)
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	seen[source] = true

	return dfs1971(source, destination, seen, graph)
}

func dfs1971(source int, destination int, seen []bool, graph [][]int) bool {
	if seen[destination] {
		return true
	}

	ans := false

	for _, v := range graph[source] {
		if !seen[v] {
			seen[v] = true
			ans = ans || dfs1971(v, destination, seen, graph)
		}
	}

	return ans
}
