func mergeAlternately(word1 string, word2 string) string {
	res := ""

	for i := 0; i < int(math.Max(float64(len(word1)), float64(len(word2)))); i++ {
		if i < len(word1) {
			res += string(word1[i])
		}
		if i < len(word2) {
			res += string(word2[i])
		}
	}

	return res
}
