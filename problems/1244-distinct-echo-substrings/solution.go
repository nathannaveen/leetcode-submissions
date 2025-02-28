func distinctEchoSubstrings(text string) int {
	res := 0

	for i := 1; i < len(text) / 2 + 1; i++ {
		m := make(map[string]int)
		str := ""
		str2 := ""
		for j := 0; j < i; j++ {
			str += string(text[j])
			str2 += string(text[j + i])
		}

		res = check(str, str2, m, res)

		for j := i; j < len(text) - i; j++ {
			str = str[1:]
			str += string(text[j])
			str2 = str2[1:]
			str2 += string(text[j+i])
			res = check(str, str2, m, res)
		}
	}

	return res
}

func check(str string, str2 string, m map[string]int, res int) int {
	if str == str2 && m[str] == 0 {
		res++
		m[str]++
		m[str2]++
	}
	return res
}