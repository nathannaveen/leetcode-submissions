func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	counter := 0
	shouldSeperate := false

	for n > 0 {
		if shouldSeperate {
			shouldSeperate = false
			res += "."
		}
		res += strconv.Itoa(n % 10)
		n /= 10
		counter++
		if counter%3 == 0 {
			shouldSeperate = true
		}

	}
	return reverseString(res)
}

func reverseString(s string) string {
	res := ""
	for _, i := range s {
		res = string(i) + res
	}
	return res
}
