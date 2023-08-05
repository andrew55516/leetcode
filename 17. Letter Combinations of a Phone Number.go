package main

var letters = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

var combinations17 []string

func letterCombinations(digits string) []string {
	combinations17 = make([]string, 0)
	comb17(digits, 0, "")
	return combinations17
}

func comb17(digits string, i int, s string) {
	if i == len(digits) {
		if len(s) > 0 {
			combinations17 = append(combinations17, s)
		}
		return
	}

	for _, v := range letters[digits[i]-'0'] {
		comb17(digits, i+1, s+string(v))
	}
}
