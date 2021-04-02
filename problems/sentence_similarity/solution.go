func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	length := len(sentence1)
	m := make(map[string][]string)

	for _, pair := range similarPairs {
		m[pair[0]] = append(m[pair[0]], pair[1])
		m[pair[1]] = append(m[pair[1]], pair[0])
	}

	for i := 0; i < length; i++ {
		if !arrayContains(m[sentence1[i]], sentence2[i]) && sentence1[i] != sentence2[i] {
			return false
		}
	}

	return true
}

func arrayContains(h []string, n string) bool {
	for _, s := range h {
		if s == n {
			return true
		}
	}
	return false
}