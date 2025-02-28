func check(nums []int) bool {
	counter := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			counter++
			if counter >= 2 {
				return false
			}
		}
		
	}

	return true
}