func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	res := 0
	for _, cost := range costs {
		if coins-cost < 0 {
			break
		}
		res++
		coins -= cost
	}
	return res
}
