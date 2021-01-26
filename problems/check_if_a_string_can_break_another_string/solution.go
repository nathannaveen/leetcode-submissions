func SortStringByCharacter(s string) []rune {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return r
}

func checkIfCanBreak(s1 string, s2 string) bool {
	one := SortStringByCharacter(s1)
	two := SortStringByCharacter(s2)
	oneGreater := -1

	if one[0] < two[0] {
		oneGreater = 2
	} else if one[0] > two[0] {
		oneGreater = 1
	} else {
		for i := 1; i < len(s2); i++ {
			if one[i] < two[i] {
				oneGreater = 2
				break
			} else if one[i] > two[i] {
				oneGreater = 1
				break
			}
		}
	}

	for i := range one {
		if oneGreater == 1 && one[i] < two[i] {
			return false
		} else if oneGreater == 2 && one[i] > two[i] {
			return false
		}
	}

	return true
}