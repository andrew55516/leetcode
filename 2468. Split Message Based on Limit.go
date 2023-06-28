package main

import (
	"fmt"
	"math"
	"strings"
)

func splitMessage(message string, limit int) []string {
	partsOrder := 0
	parts := make([]string, 0)

	for i := 1; i <= 6; i++ {
		maxMessageLen := 0
		for j := 1; j <= i; j++ {
			maxMessageLen += (limit - (3 + j + i)) * (pow(10, j) - pow(10, j-1))
		}
		if len(message) <= maxMessageLen {
			partsOrder = i
			break
		}
	}

	if partsOrder == 0 {
		return parts
	}

	builder := strings.Builder{}

	k := 0
	current := 0
	for _, j := range message {
		builder.WriteRune(j)
		current += 1
		if current == limit-(3+len(fmt.Sprintf("%d", k+1))+partsOrder) {
			builder.WriteString(fmt.Sprintf("<%d/", k+1))
			parts = append(parts, builder.String())
			builder.Reset()
			current = 0
			k += 1
		}
	}

	if current > 0 {
		builder.WriteString(fmt.Sprintf("<%d/", k+1))
		parts = append(parts, builder.String())
		k += 1
	}

	for i, _ := range parts {
		parts[i] = fmt.Sprintf("%s%d>", parts[i], k)
	}

	return parts
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
