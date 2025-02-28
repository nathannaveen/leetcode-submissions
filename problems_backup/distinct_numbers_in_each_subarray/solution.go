func distinctNumbers(nums []int, k int) []int {
	counter := 0
	resCounter := 0
	m := make(map[int]int)
	res := make([]int, len(nums)-k+1)

	for i := 0; i < k; i++ {
		m[nums[i]]++
		if m[nums[i]] == 1 {
			counter++
		}
	}
	res[resCounter] = counter

	for i := k; i < len(nums); i++ {
		m[nums[i-k]]--
		m[nums[i]]++
		if m[nums[i-k]] == 0 {
			counter--
		}
		if m[nums[i]] == 1 && nums[i] != nums[i - k] {
			counter++
		}
		resCounter++
		res[resCounter] = counter
	}
	return res
}