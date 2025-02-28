func getMaximumGenerated(n int) int {
	max := float64(1)
	arr := make([]float64, n+1)
	counter := 2
	if n == 0 {
		return 0
	}
	arr[0] = 0
	arr[1] = 1

	for i := 1; i < (n-1)/2+1; i++ {
		arr[counter] = arr[i]
		arr[counter+1] = arr[i] + arr[i+1]
		max = math.Max(arr[counter+1], max)
		counter += 2
	}

	if n%2 == 0 {
		arr[counter] = arr[(n-1)/2]
		max = math.Max(max, arr[(n-1)/2]+arr[(n-1)/2+1])
	}

	return int(max)
}