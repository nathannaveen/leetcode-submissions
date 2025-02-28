type point struct {
	x int
	y int
}

func minTotalDistance(grid [][]int) int {
	ones := []point{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ones = append(ones, point{x: j, y: i})
			}
		}
	}

	minimum := 2147483648
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			minimum = min(manhattanDistance(i, j, ones), minimum)
		}
	}
	return minimum
}

func manhattanDistance(i, j int, ones []point) int {
	sum := 0
	for _, thePoint := range ones {
		sum += abs(j-thePoint.x) + abs(i-thePoint.y)
	}
	return sum
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}