func longestWord(words []string) string {
	sort.Strings(words)
	g := make(map[string]int)
	max := ""

	for _, word := range words {
		if g[word[:len(word)-1]] != 0 || len(word) == 1 {
			if len(word) > len(max) {
				max = word
			}
			g[word] = 1
		}
	}

	return max
}