func findMissingRanges(nums []int, lower int, upper int) []string {

	if len(nums) == 0 {
		if lower == upper {
			return []string{strconv.Itoa(lower)}
		} else {
			return []string{strconv.Itoa(lower) + "->" + strconv.Itoa(upper)}
		}
	}
	res := []string{}

	difference := nums[0] - lower
	if abs(difference) == 1 {
		res = append(res, strconv.Itoa(lower))
	} else if abs(difference) > 1 {
		res = append(res, strconv.Itoa(lower)+"->"+strconv.Itoa(nums[0]-1))
	}

	for i := 1; i < len(nums); i++ {
		previousNumber := nums[i-1]
		currentNumber := nums[i]

		difference2 := currentNumber - previousNumber
		if difference2 == 2 {
			res = append(res, strconv.Itoa(previousNumber+1))
		} else if difference2 > 2 {
			res = append(res, strconv.Itoa(previousNumber+1)+"->"+strconv.Itoa(currentNumber-1))
		}
	}

	lastNumber := nums[len(nums)-1]

	difference3 := upper - lastNumber
	if abs(difference3) == 1 {
		res = append(res, strconv.Itoa(upper))
	} else if difference3 != 0 {
		res = append(res, strconv.Itoa(lastNumber+1)+"->"+strconv.Itoa(upper))
	}

	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}