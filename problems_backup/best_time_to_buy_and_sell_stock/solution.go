func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	min := prices[0]
	max := 0

	for _,price := range prices {
		if price < min {
			min = price
		} else if (price-min) > max{
			max = price-min
		}
	}

	return  max
}