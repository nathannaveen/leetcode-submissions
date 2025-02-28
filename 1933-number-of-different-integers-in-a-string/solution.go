func numDifferentIntegers(word string) int {
	m := make(map[string] bool)
	number := ""
	hasZeros := false

	for _, c := range word {
		if c == '0' && len(number) == 0 { // is digit
			hasZeros = true
		} else if c - ':' < 0 {
			number += string(c)
		} else if c - ':' > 0 { // is letter
			addNumberToMap(number, hasZeros, m)
			hasZeros = false
			number = ""
		}
	}

	addNumberToMap(number, hasZeros, m)

	return len(m)
}

func addNumberToMap(number string, hasZeros bool, m map[string] bool) {
	if len(number) == 0 && hasZeros {
		m["0"] = true
	} else if len(number) != 0 {
		m[number] = true
	}
}