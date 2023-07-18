package main

import "strings"

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, p := range parts {
		if p == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if p != "." && p != "" {
			stack = append(stack, p)
		}
	}

	return "/" + strings.Join(stack, "/")
}
