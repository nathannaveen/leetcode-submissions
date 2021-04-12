func minDeletions(s string) int {
	res := 0
	arr := make([]int, 100000)
	m := make(map[int32]int)
	goods := []int{}

	for _, i2 := range s {
		m[i2]++
	}

	for _, i := range m {
		arr[i]++
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] == 0 {
			goods = append(goods, i)
		} else if arr[i] != 1 {
			for arr[i] > 1 {
				if len(goods) == 0 {
					res += i
					arr[i] -= 1
				} else {
					res += i - goods[len(goods)-1]
					arr[i] -= 1
					goods = goods[:len(goods)-1]
				}
			}
		}
	}

	return res
}
