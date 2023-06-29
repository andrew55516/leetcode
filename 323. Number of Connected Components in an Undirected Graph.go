package main

func countComponents(n int, edges [][]int) int {
	graph := make([][]int, n)
	seen := make([]bool, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	ans := 0

	for i := 0; i < n; i++ {
		if !seen[i] {
			ans++
			seen[i] = true
			dfs323(graph, seen, i)
		}
	}

	return ans
}

func dfs323(graph [][]int, seen []bool, node int) {
	for _, v := range graph[node] {
		if !seen[v] {
			seen[v] = true
			dfs323(graph, seen, v)
		}
	}
}
