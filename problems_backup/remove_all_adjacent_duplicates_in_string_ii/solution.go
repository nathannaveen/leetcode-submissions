func removeDuplicates(s string, k int) string {
	h := []string{}
	for _, i2 := range s {
		h = append(h, string(i2))
		if len(h) >= k {
			allEqualsCharacter := true
			for i := 0; i < k; i++ {
				if h[len(h)-(i+1)] != string(i2) {
					allEqualsCharacter = false
					break
				}
			}
			if allEqualsCharacter {
				h = h[:len(h)-k]
			}
		}
	}
	m := ""
	for _, s2 := range h {
		m += s2
	}
	return m
}