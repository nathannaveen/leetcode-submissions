func twoSum(numbers []int, target int) []int {
	h := []int{}
	left := 0
	right := len(numbers) - 1
	sort.Ints(numbers)

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum < target {
			left ++
		} else if sum > target {
			right --
		} else {
			break
		}

	}

	h = append(h, left + 1, right + 1)
	return h
}
