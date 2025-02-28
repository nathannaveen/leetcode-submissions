func nextGreaterElements(nums []int) []int {
	h := []int{}
	for i := 0; i < len(nums); i++ {
		counter := 0
		j := i + 1
		for true {
			if j == len(nums) {
				j = 0
			}
			if nums[j] > nums[i] {
				h = append(h, nums[j])
				break
			}
			if counter == len(nums){
                h = append(h, -1)
				break
			}
			j ++
			counter ++
		}
	}
	return h
}