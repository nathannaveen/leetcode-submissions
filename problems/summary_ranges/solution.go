func summaryRanges(nums []int) []string {
	h := []string{}
	if len(nums) == 0 {
		return h
	}
	first := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			s := strconv.Itoa(first)
			if nums[i-1] != first {
				s += "->" + strconv.Itoa(nums[i-1])
			}
			h = append(h, s)
			first = nums[i]
		}
	}

	s := strconv.Itoa(first)
	if nums[len(nums)-1] != first {
		s += "->" + strconv.Itoa(nums[len(nums)-1])
	}
	h = append(h, s)

	return h
}
