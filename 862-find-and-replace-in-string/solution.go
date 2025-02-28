func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	res := ""

	for i := 0; i < len(S); i++ {
		index := 0 // this is the index that is used for indexes, sources, and targets

		for i2, i3 := range indexes {
			if i3 == i {
				index = i2 // we do this because the indexes are not sorted
			}
		}

		if index < len(indexes) && i == indexes[index] {
			currentString := sources[index]
			isSubstring := true
			for i2 := range currentString { // checking whether the whole substring is equal
				if string(S[i+i2]) != string(currentString[i2]) { 
					isSubstring = false
					break
				}
			}

			if isSubstring {
				res += targets[index]
				i += len(currentString) - 1 // we have to add to i because we dont want to repeat characters
			} else {
				res += string(S[i]) // if there is not a whole substring just add the character
			}
			index++
		} else {
			res += string(S[i]) // if there is no index in indexes that works for the current index
		}
	}

	return res
}
