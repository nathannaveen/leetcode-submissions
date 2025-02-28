func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		res += int(num) & 1
		num >>= 1
	}
	return res
}