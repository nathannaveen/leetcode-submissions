func leastBricks(wall [][]int) int {
	m := make(map[int] int)
	
	for _, row := range wall {
		total := 0
		for _, brick := range row[:len(row)-1] {
			total, m[total+brick] = total+brick, m[total+brick]+1
		}
	}
	maximum := 0
	for _, i := range m {
		if i > maximum {
			maximum = i
		}
	}
	return len(wall) - maximum
}
