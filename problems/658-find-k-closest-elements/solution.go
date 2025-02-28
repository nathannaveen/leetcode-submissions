func findClosestElements(arr []int, k int, x int) []int {
	for i := 1; i < len(arr); i++ {
		if i >= 1 && abs(arr[i-1]-x) > abs(arr[i]-x) {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i -= 2
		}
	}
	arr = arr[:k]

	for i := 1; i < len(arr); i++ {
		if i >= 1 && arr[i-1] > arr[i] {
			arr[i - 1], arr[i] = arr[i], arr[i - 1]
            i -= 2
		}
	}
	
	return arr
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}