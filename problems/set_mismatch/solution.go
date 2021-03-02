func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	double, skip := 0, 0
	temp := make([]int, 10001)

	for _, num := range nums {
		temp[num]++
	}

	for i, value := range temp {
		if value == 2 {
			double = i
			if skip != 0 {
				return []int{double, skip}
			}
		} else if value == 0 {
			skip = i
			if double != 0 {
				return []int{double, skip}
			}
		}
	}
	
	return []int{double, double + 1}
}
