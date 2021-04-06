func moveZeroes(nums []int) {
	counter := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[counter], nums[i] = nums[i], nums[counter]
			counter++
		}
	}
}