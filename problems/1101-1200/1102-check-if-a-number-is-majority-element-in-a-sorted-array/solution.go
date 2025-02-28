func isMajorityElement(nums []int, target int) bool {
	m := make(map[int] int)

	for _, num := range nums {
		m[num]++
	}
	return m[target] > len(nums) / 2
}