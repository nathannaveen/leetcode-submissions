
func licenseKeyFormatting(S string, K int) string {
	str := strings.ReplaceAll(S, "-", "")
	str = strings.ToUpper(str)
	str = flipAndAddDashes(str, K)
	res := ""
	chars := []rune(str)

	for i := len(str) - 1; i >= 0; i-- {
		res += string(chars[i])
	}

	return res
}

func flipAndAddDashes(s string, k int) string {
	result := ""
	characters := []rune(s)
	counter := 1
	for i := len(s) - 1; i >= 0; i-- {
		result += string(characters[i])
		if counter%k == 0 && counter != len(s) {
			result += "-"
		}
		counter++
	}
	return result
}