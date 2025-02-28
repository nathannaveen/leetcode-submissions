func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	res := []int{}
	a, b, c := 0, 0, 0

	for a < len(arr1) && b < len(arr2) && c < len(arr3) {
		maximum := max(arr1[a], max(arr2[b], arr3[c]))
		if arr1[a] == maximum && arr2[b] == maximum && arr3[c] == maximum {
			res = append(res, arr1[a])
			a, b, c = a+1, b+1, c+1
		} else {
			if arr1[a] != maximum {
				a++
			}
			if arr2[b] != maximum {
				b++
			}
			if arr3[c] != maximum {
				c++
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}