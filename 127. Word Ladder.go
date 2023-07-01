package main

import "strings"

type state struct {
	val   []string
	steps int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[int][]string, len(wordList))
	for i, s := range wordList {
		if s != beginWord {
			wordMap[i] = strings.Split(s, "")
		}
	}

	queue := make([]state, 0)
	queue = append(queue, state{strings.Split(beginWord, ""), 1})

	for len(queue) > 0 {
		s := queue[0]
		if strings.Join(s.val, "") == endWord {
			return s.steps
		}
		queue = queue[1:]
		
		for k, v := range wordMap {
			d := 0
			for k := 0; k < len(s.val); k++ {
				if s.val[k] != v[k] {
					d++
				}
			}
			if d == 1 {
				queue = append(queue, state{v, s.steps + 1})
				delete(wordMap, k)
			}
		}
	}

	return 0
}
