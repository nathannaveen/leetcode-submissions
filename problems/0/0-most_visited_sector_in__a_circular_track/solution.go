func mostVisited(n int, rounds []int) []int {
	h := make([]int, n)
	res := []int{}
	max := 0
	for i := 1; i < len(rounds); i++ {
		end := rounds[i]
		if rounds[i-1] > rounds[i] {
			end = rounds[i] + n
		}
		for j := rounds[i-1]; j < end; j++ {
			if j%n == 0 {
				h[len(h)-1]++
			} else {
				h[(j%n)-1]++
			}
		}
	}
	h[rounds[len(rounds) - 1] - 1]++
	for i, i2 := range h {
		if i2 > max {
			max = i2
			res = []int{i + 1}
		} else if i2 == max {
			res = append(res, i+1)
		}
	}
	return res
}