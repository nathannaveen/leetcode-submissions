func twoSum(numbers []int, target int) []int {
	 h :=[]int{}
	left := 0
	right := len(numbers) - 1;
	sort.Ints(numbers)

	for i := 0; i < len(numbers); i++ {
		if numbers[left] + numbers[right] < target {
			left ++
		}
		if numbers[left] + numbers[right] > target {
			right --
		}
		if numbers[left] + numbers[right] == target {
			h = append(h, left + 1)
			h = append(h, right + 1)
			return h
		}
	}
	return h
}