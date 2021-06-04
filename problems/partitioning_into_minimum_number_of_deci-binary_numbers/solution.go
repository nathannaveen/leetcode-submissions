func minPartitions(n string) int {
	max := '0'

	for _, i := range n {
		if i > max {
			max = i
		}
	}
	
	return int(max - '0')
}