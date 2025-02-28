func singleNumber(nums []int) []int {
	res := []int{}
	sort.Ints(nums)
	for i := 1; i < len(nums) - 1; i++ {
		if i != 0 && i != len(nums)-1 && nums[i] != nums[i-1] && nums[i+1] != nums[i] {
			res = append(res, nums[i])
		}	
	}
	
	if nums[0] != nums[1] {
		res = append(res, nums[0])
	}
	if nums[len(nums)-1] != nums[len(nums)-2] {
		res = append(res, nums[len(nums)-1])
	}

	return res
}
