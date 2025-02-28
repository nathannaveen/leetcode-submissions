func differByOne(dict []string) bool {
	sort.Strings(dict)
	for i := 0; i < len(dict); i++ {
		for j := i + 1; j < len(dict); j++ {
			if differByOneCharacter(dict[i], dict[j]) {
				return true
			}
		}
	}
	return false
}

func differByOneCharacter(a, b string) bool {
	diffCounter := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] { // We can do this because in the constrains dict[i].length == dict[j].length
			diffCounter++
		}
	}
	return diffCounter == 1 // All strings are unique so diffCounter != 0
}