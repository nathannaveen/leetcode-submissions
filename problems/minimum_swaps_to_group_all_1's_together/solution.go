func minSwaps(data []int) int {
	oneCounter, maximum, counter := 0, 0, 0

	for _, i := range data {
		oneCounter += i
	}

	for i := 0; i < oneCounter; i++ {
		counter += data[i]
	}
	maximum = max(maximum, counter)
	
	for i := oneCounter; i < len(data); i++ {
		counter += data[i] - data[i - oneCounter]
		maximum = max(maximum, counter)
	}

	return oneCounter - maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}