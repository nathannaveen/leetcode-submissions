func findMinFibonacciNumbers(k int) int {
	res, a, b := 0, 1, 1
	arr := []int{1}
	for b <= k {
		a, b = b, a+b
		arr = append(arr, b)
	}
	index := len(arr) - 1
	for k > 0 {
		if k >= arr[index] {
			res++
			k -= arr[index]
			continue
		}
		index--
	}

	return res
}