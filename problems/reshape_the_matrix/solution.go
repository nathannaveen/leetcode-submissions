func matrixReshape(nums [][]int, r int, c int) [][]int {

	if len(nums) * len(nums[0]) != r * c {
		return nums
	}

	res := make([][]int, r)
	rCounter := 0
	cCounter := 0

	for _, num := range nums {
		for _, i3 := range num {
			res[rCounter] = append(res[rCounter], i3)
			cCounter++
			if cCounter == c {
				cCounter = 0
				rCounter++
				if rCounter == r {
					return res
				}
			}
		}
	}
	return res
}