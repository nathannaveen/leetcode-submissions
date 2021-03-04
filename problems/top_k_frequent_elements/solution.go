func topKFrequent(nums []int, k int) []int {
	arr := [][]int{}

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
		if i >= 1 && arr[i-1][0] < arr[i][0] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i -= 2
		}
	}

	nums = []int{}
	for i := 0; i < k; i++ {
		nums = append(nums, arr[i][1])
	}

	return nums
}
