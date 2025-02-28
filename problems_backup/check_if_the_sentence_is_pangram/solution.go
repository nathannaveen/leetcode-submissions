func checkIfPangram(sentence string) bool {
	arr := make([]int, 26) 
	counter := 0

	for _, i := range sentence {
		arr[int(i-'a')]++
		if arr[int(i-'a')] == 1 {
			counter++
		}
	}

	return counter == 26
}