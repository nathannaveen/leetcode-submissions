func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0

	for left < right {
		max = int(math.Max(float64(max),float64((right - left) * int(math.Min(float64(height[left]), float64(height[right]))))))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
