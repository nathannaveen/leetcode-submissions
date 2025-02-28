func dominantIndex(nums []int) int {
	h := []int{}
	for _, num := range nums {
		h = append(h, num)
	}
	sort.Ints(h)
	if len(nums) > 1 && h[len(h)-1] >= h[len(h)-2]*2 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == h[len(h) - 1] {
				return i
			}
		}
	} else if len(nums) == 1 {
		return 0
	}
	return -1
}