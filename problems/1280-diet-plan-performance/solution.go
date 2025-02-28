func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	sum := 0
	res := 0
	for i := 0; i < k; i++ {
		sum += calories[i]
	}
	res += comparison(sum, lower, upper)

	for i := k; i < len(calories); i++ {
		sum -= calories[i - k] - calories[i]
		res += comparison(sum, lower, upper)
	}
	return res
}

func comparison(sum int, lower int, upper int) int {
	if sum < lower {
		return -1
	} else if sum > upper {
		return 1
	}
	return 0
}