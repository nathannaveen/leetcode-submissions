func lexicalOrder(n int) []int {

	strs := []string{}
	res := []int{}

	for i := 1; i <= n; i++ {
		strs = append(strs, strconv.Itoa(i))
	}
	sort.Strings(strs)
	for i := range strs {
		n, _ = strconv.Atoi(strs[i])
		res = append(res, n)
	}
	return res
}