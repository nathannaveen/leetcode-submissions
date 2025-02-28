func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int)
	res := []string{}
	minimum := 2001

	for i, s := range list1 {
		m[s] = i
	}

	for i, s := range list2 {
		if _, ok := m[s]; ok {
			if m[s] + i < minimum {
				minimum = m[s] + i
				res = []string{s}
			} else if m[s] + i == minimum {
				res = append(res, s)
			}
		}
	}
	return res
}