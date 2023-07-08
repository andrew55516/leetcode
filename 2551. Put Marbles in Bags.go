package main

import "sort"

func putMarbles(weights []int, k int) int64 {
	splitWeights := make([]int, 0, len(weights)-1)
	for i := 0; i < len(weights)-1; i++ {
		splitWeights = append(splitWeights, weights[i]+weights[i+1])
	}

	sort.Ints(splitWeights)
	max := weights[0] + weights[len(weights)-1]
	min := weights[0] + weights[len(weights)-1]
	for i := 0; i < k-1; i++ {
		min += splitWeights[i]
	}
	for i := len(splitWeights) - k + 1; i < len(splitWeights); i++ {
		max += splitWeights[i]
	}
	return int64(max - min)
}
