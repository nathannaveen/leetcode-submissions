func toGoatLatin(S string) string {
	h := strings.Split(S, " ")
	vowels := make(map[uint8]int)
	vowels['a'] = 1; vowels['e'] = 1; vowels['i'] = 1; vowels['o'] = 1; vowels['u'] = 1
	vowels['A'] = 1; vowels['E'] = 1; vowels['I'] = 1; vowels['O'] = 1; vowels['U'] = 1
	a := "a"
	for i := 0; i < len(h); i++ {
		if vowels[h[i][0]] == 0 {
			h[i] = h[i][1:] + string(h[i][0])
		}
		h[i] += "ma" + a
		a += "a"
	}
	return strings.Join(h, " ")
}