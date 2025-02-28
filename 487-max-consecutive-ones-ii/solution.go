func findMaxConsecutiveOnes(nums []int) int {
	zeros := []int{}
	maximum := 0
	zeroCounter := 0
	counter := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if zeroCounter == 0 {
				zeroCounter++
				counter++
			} else {
				zeroCounter = 0
				maximum = max(maximum, counter)
				counter = 0
				if zeros[len(zeros)-1]+1 < len(nums) {
					i = zeros[len(zeros)-1] + 1
					if nums[zeros[len(zeros)-1] + 1] == 1 {
						counter = 1
					}
					zeros = zeros[:len(zeros)-1]
				}
			}
			zeros = append(zeros, i)
		} else {
			counter++
		}
	}
	maximum = max(maximum, counter)
	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}