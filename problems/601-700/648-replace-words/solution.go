func replaceWords(dictionary []string, sentence string) string {
	res := ""
	words := strings.Split(sentence, " ")
	sort.Strings(dictionary)
	for i, word := range words {
		prefix := false
		for _, s := range dictionary {
			if strings.HasPrefix(word, s) {
				res += s
				prefix = true
				break
			}
		}
		if !prefix {
			res += word
		}
		if i != len(words)-1 {
			res += " "
		}
	}

	return res
}
