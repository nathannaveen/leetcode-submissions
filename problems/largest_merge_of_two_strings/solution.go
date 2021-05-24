func largestMerge(word1 string, word2 string) string {
	res := ""

	for len(word1) != 0 && len(word2) != 0 {
		if word1 > word2 {
			res += string(word1[0])
			word1 = word1[1:]
		} else {
			res += string(word2[0])
			word2 = word2[1:]
		}
	}
	res += word2
	res += word1
	return res
}