func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	shouldAdd := true
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				shouldAdd = false
				break
			}
		}
		if shouldAdd {res++} else {shouldAdd = true}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
