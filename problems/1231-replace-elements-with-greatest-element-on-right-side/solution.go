func replaceElements(arr []int) []int {
	for i := range arr{
		max := 0
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		arr[i] = max
	}
	arr[len(arr) - 1] = -1
	return arr
}