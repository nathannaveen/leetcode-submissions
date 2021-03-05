func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	index := 0
	if ruleKey == "color" {
		index = 1
	} else if ruleKey == "name" {
		index = 2
	}

	for _, item := range items {
		if item[index] == ruleValue {
			res++
		}
	}
	
	
	return res
}