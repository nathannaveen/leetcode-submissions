func numOfSubarrays(arr []int, k int, threshold int) int {
	sum := 0
	res := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if float64(sum)/float64(k) >= float64(threshold) {
		res++
	}

	for i := k; i < len(arr); i++ {
		sum += arr[i] - arr[i-k]
		if float64(sum)/float64(k) >= float64(threshold) {
			res++
		}
	}
	return res
}
