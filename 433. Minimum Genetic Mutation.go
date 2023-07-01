package main

import "strings"

type gene struct {
	code  []string
	steps int
}

func minMutation(startGene string, endGene string, bank []string) int {
	choice := []string{"A", "C", "G", "T"}
	start := strings.Split(startGene, "")
	bankMap := make(map[string]struct{}, len(bank))
	for _, s := range bank {
		bankMap[s] = struct{}{}
	}

	seen := make(map[string]struct{}, 0)
	seen[startGene] = struct{}{}
	queue := make([]gene, 0)
	queue = append(queue, gene{code: start, steps: 0})

	for len(queue) > 0 {
		g := queue[0]

		if strings.Join(g.code, "") == endGene {
			return g.steps
		}

		queue = queue[1:]

		for _, c := range choice {
			for i := 0; i < len(g.code); i++ {
				if g.code[i] == c {
					continue
				}
				next := gene{
					code:  make([]string, len(g.code)),
					steps: g.steps + 1,
				}
				copy(next.code, g.code)

				next.code[i] = c
				nextStr := strings.Join(next.code, "")
				if _, ok := bankMap[nextStr]; ok {
					if _, ok := seen[nextStr]; !ok {
						queue = append(queue, next)
						seen[nextStr] = struct{}{}
					}
				}
			}
		}

	}

	return -1
}
