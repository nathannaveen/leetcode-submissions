func findBuildings(heights []int) []int {
	maximum := 0
	arr := []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		if heights[i] > maximum {
			maximum = heights[i]
			arr = append(arr, i)
		}
	}
	return reverseArr(arr)
}

func reverseArr(arr []int) []int {
	left, right := 0, len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left, right = left + 1, right - 1
	}
	return arr
}