func frequencySort(nums []int) []int {
	arr := [][]int{}
	counter := 0

	for _, num := range nums {
		contains := false

		for _, ints := range arr {
			if ints[1] == num {
				ints[0]++
				contains = true
				break
			}
		}

		if !contains {
			arr = append(arr, []int{1, num})
		}
	}

	for i := 1; i < len(arr); i++ {
		if i >= 1 && ((arr[i-1][0] > arr[i][0]) || (arr[i-1][0] == arr[i][0] && arr[i-1][1] < arr[i][1])) {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i -= 2
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < arr[i][0]; j++ {
			nums[counter] = arr[i][1]
			counter++
		}
	}

	return nums
}
