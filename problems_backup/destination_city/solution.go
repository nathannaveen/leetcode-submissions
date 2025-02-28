func destCity(paths [][]string) string {
	start := make(map[string]int)
	destination := make(map[string]int)

	for _, path := range paths {
		start[path[0]] = 1
		destination[path[1]] = 1
		
		if destination[path[0]] == 1 {
			delete(destination, path[0])
			delete(start, path[0])
		}
		if start[path[1]] == 1 {
			delete(start, path[1])
			delete(destination, path[1])
		}
	}
	
	for s, _ := range destination {
		return s
	}
	return ""
}