package main

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		dict[word] = struct{}{}
	}
	return bfs139(s, dict)
}

func bfs139(s string, dict map[string]struct{}) bool {
	queue := make([]int, 0)
	queue = append(queue, 0)

	seen := make(map[int]struct{}, 0)

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		for end := start + 1; end <= start+20 && end <= len(s); end++ {
			if _, ok := seen[end]; ok {
				continue
			}
			if _, ok := dict[s[start:end]]; ok {
				if end == len(s) {
					return true
				}
				queue = append(queue, end)
				seen[end] = struct{}{}
			}
		}
	}
	return false
}
