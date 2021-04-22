func buddyStrings(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[uint8]int)
	temp := -1
	counter := 0

	for i := 0; i < len(a); i++ {
		m[a[i]]++
		if a == b && m[a[i]] == 2 {
			return true
		}
		if a[i] != b[i] {
			counter++
			if temp == -1 {
				temp = i
			} else if a[temp] != b[i] || a[i] != b[temp] {
				return false
			}
		}
	}

	return counter == 2
}