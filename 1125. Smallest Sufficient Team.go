package main

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	m := make(map[string]int, len(req_skills))
	for i, s := range req_skills {
		m[s] = i
	}
	personSkills := make([]int, len(people))
	for i, v := range people {
		for _, s := range v {
			personSkills[i] |= 1 << m[s]
		}
	}

	skillsMask := 0
	for _ = range req_skills {
		skillsMask = (skillsMask << 1) | 1
	}

	dp := make([]int, skillsMask+1)
	for i := range dp {
		dp[i] = 1<<len(people) - 1
	}
	dp[0] = 0
	for i := 1; i <= skillsMask; i++ {
		for j := 0; j < len(people); j++ {
			mask := i &^ personSkills[j]
			if mask != i {
				team := dp[mask] | (1 << j)
				if len(toArray1125(team)) < len(toArray1125(dp[i])) {
					dp[i] = team
				}
			}
		}
	}
	return toArray1125(dp[skillsMask])
}

func toArray1125(n int) []int {
	i := 0
	ans := make([]int, 0)
	for n > 0 {
		if n%2 == 1 {
			ans = append(ans, i)
		}
		n /= 2
		i++
	}
	return ans
}
