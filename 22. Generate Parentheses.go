package main

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	backtracking22(n, 0, 0, "", &ans)
	return ans
}

func backtracking22(n, opened, closed int, path string, combinations *[]string) {
	if opened == closed && opened == n {
		*combinations = append(*combinations, path)
		return
	}

	if opened < n {
		backtracking22(n, opened+1, closed, path+"(", combinations)
	}

	if opened > closed {
		backtracking22(n, opened, closed+1, path+")", combinations)
	}
}
