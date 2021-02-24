func arrangeWords(text string) string {
	split := strings.Split(text, " ")

	for i := 1; i < len(split); i++ {
		if len(split[i-1]) > len(split[i]) {
			hasPut := false
			l := split[i]
			for j := i - 1; j >= 0; j-- {
				if (len(split[j]) <= len(l)) || (j >= 1 && len(split[j-1]) <= len(l) && len(split[j]) >= len(l)) {
					split[j+1] = split[j]
					split[j] = l
					hasPut = true
					break
				}
				split[j+1] = split[j]
			}

			if !hasPut {
				split[0] = l
			}
		}
	}

	g := strings.Join(split, " ")
	g = strings.ToLower(g)
	res := strings.ToUpper(string(g[0]))
	res += g[1:]

	return res
}
