func truncateSentence(s string, k int) string {
	arr := strings.Split(s, " ")[:k]
	return strings.Join(arr, " ")
}