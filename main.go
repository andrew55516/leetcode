package main

import "fmt"

func main() {
	req := []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"}
	people := [][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}}
	fmt.Println(SmallestSufficientTeam(req, people))
}
