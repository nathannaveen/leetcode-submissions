func rotate(nums []int, k int) {
	h := make([]int, len(nums))

	for i, num := range nums {
		h[i] = num
	}

	for i, i2 := range h {
		g := i + k
		if i+k > len(h)-1 {
			g = (i + k) % len(h)
		}
		nums[g] = i2
	}
}
