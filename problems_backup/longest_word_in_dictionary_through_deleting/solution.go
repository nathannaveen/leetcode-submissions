func findLongestWord(s string, d []string) string {
	res := ""
	for _, word := range d {
		i := 0
		for j := 0; j < len(s) && i < len(word); j++ {
			if word[i] == s[j] {
				i++
			}
		}
		if i == len(word) {
			if len(word) > len(res) || len(word) == len(res) && word < res {
				res = word
			}
		}
	}
	return res
}
