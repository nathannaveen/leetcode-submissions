func maxNumberOfApples(arr []int) int {

	for i := 1; i < len(arr); i++ {
		if i >= 1 && arr[i - 1] > arr[i] {
			arr[i - 1], arr[i] = arr[i], arr[i - 1]
			i -= 2
		}
	}

	sum := 0
	for i, apple := range arr {
		if sum + apple > 5000 {
			return i 
		}
		sum += apple
	}
	return len(arr)
}