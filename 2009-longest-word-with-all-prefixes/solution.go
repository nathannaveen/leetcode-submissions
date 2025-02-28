func longestWord(words []string) string {
	m := make(map[string]int)
	maximum := ""
	sort.Strings(words)
    fmt.Println(words)
	for _, word := range words {
		if len(word) == 1 || m[word[:len(word) - 1]] >= 1 {
			m[word]++
			if len(word) > len(maximum) {
				maximum = word
			}
		}
	}
	return maximum
}