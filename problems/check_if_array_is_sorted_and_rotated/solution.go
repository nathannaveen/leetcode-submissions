func check(nums []int) bool {
	
	h := []int{}

	for _, num := range nums {
		h = append(h, num)
	}

	sort.Ints(h)

	for i := 0; i < len(nums); i++ {
		if reflect.DeepEqual(h, nums) {
			return true
		}
		
		n := nums[0]

		for j := 1; j < len(nums); j++ {
			nums[j - 1] = nums[j]
		}
		nums[len(nums) - 1] = n
	}
	return false
}