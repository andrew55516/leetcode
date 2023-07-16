package main

func eventualSafeNodes(graph [][]int) []int {
	//graphMap := make(map[int][]int)
	//for i, v := range graph {
	//	graphMap[i] = v
	//}
	seen := make([]bool, len(graph))
	safe := make([]bool, len(graph))

	for i, v := range graph {
		if len(v) == 0 {
			safe[i] = true
		}
	}

	for i, _ := range graph {
		if !seen[i] {
			seen[i] = true
			path := []int{i}
			safe[i] = dfs802(graph, seen, safe, &path, i)
		}
	}
	ans := make([]int, 0)
	for i, v := range safe {
		if v {
			ans = append(ans, i)
		}
	}
	return ans
}

func dfs802(graph [][]int, seen, safe []bool, path *[]int, node int) bool {
	if safe[node] {
		return true
	}
	ans := true
	for _, v := range graph[node] {
		if !seen[v] {
			seen[v] = true
			*path = append(*path, v)
			ok := dfs802(graph, seen, safe, path, v)
			if ok {
				safe[v] = true
			}
			ans = ans && ok
		} else if !safe[v] {
			return false
		}
	}

	return ans
}
