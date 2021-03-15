func topKFrequent(words []string, k int) []string {
	frequency := [][]int{}

	for i, word := range words {
		contains := false

		for _, ints := range frequency {
			if words[ints[0]] == word {
				contains = true
				ints[1]++
				break
			}
		}
		if !contains {
			frequency = append(frequency, []int{i, 1})
		}
	}

	for i := 0; i < len(frequency); i++ {
		if i >= 1 && (frequency[i][1] < frequency[i-1][1] ||
			(frequency[i][1] == frequency[i-1][1] && words[frequency[i][0]] > words[frequency[i-1][0]])) {

			frequency[i], frequency[i-1] = frequency[i-1], frequency[i]
			i -= 2
		}
	}
	res := []string{}

	for i := len(frequency) - 1; i >= len(frequency) - k; i-- {
		res = append(res, words[frequency[i][0]])
	}
	return res
}
