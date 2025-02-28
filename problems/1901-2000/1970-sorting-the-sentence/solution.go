func sortSentence(s string) string {
	arr := strings.Split(s, " ")
	res := make([]string, len(arr))

	for i := 0; i < len(arr); i++ {
		res[arr[i][len(arr[i]) - 1] - '1'] = arr[i][:len(arr[i]) - 1]
	}
	return strings.Join(res, " ")
}
