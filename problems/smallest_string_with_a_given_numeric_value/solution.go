func getSmallestString(n int, k int) string {
	k -= n
	res := make([]string, n)
    
	for i := len(res) - 1; i >= 0; i-- {
		if k == 0 {
			res[i] = "a"
		} else if k <= 25 {
			res[i] = string(k + 'a')
			k = 0
		} else {
			res[i] = "z"
			k -= 25
		}
	}

	return strings.Join(res, "")
}