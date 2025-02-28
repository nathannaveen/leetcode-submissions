func sumOfUnique(nums []int) int {
	sum := 0
	m := make(map[int] int)
	for _, num := range nums {
		m[num]++
	}
	for i, i2 := range m {
		if i2 == 1 {
			sum += i
		}
	}
	return sum
}
