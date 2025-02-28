func missingNumber(arr []int) int {
	diff1 := arr[1] - arr[0]
	diff2 := arr[2] - arr[1]
	if len(arr) == 3 {
		if diff1 > diff2 && diff2 < 0 {
			return arr[1] + diff1
		} else if diff2 < 0 {
			return arr[0] + diff2
		}
		if diff1 > diff2 {
			return arr[0] + diff2
		} else {
			return arr[1] + diff1
		}
	}
	diff3 := arr[3] - arr[2]
	diff := 0

	if diff1 == diff2 {
		diff = diff1
	} else if diff2 == diff3 {
		diff = diff2
	} else if diff3 == diff1 {
		diff = diff3
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != diff {
			return arr[i] + diff
		}
	}
	return arr[0]
}