type points struct {
	x, y int
}

type key struct {
	val int
	p   []points
}

func containsCycle(grid [][]byte) bool {
	arr := [][]key{}

	for i := 0; i < len(grid); i++ {
		arr = append(arr, make([]key, len(grid[0])))
		for j := 0; j < len(grid[0]); j++ {
			arr[i][j].val = 1
			if i > 0 && grid[i - 1][j] == grid[i][j] {
				arr[i][j].val *= int(grid[i - 1][j] - 'a') + 2
				arr[i][j].p = append(arr[i][j].p, arr[i - 1][j].p...)
			}
			if j > 0 && grid[i][j - 1] == grid[i][j] {
				arr[i][j].val *= int(grid[i][j - 1] - 'a') + 2
				arr[i][j].p = append(arr[i][j].p, arr[i][j - 1].p...)
			}
			if len(arr[i][j].p) == 0 {
				a := points{j, i}
				arr[i][j].p = []points{ a }
			}

			if arr[i][j].val != 1 && arr[i][j].val != int(grid[i][j] - 'a') + 2 {

				b, done := ifStartIsSame(arr, i, j)
				if done {
					return b
				}
			}
		}
	}

	return false
}

func ifStartIsSame(arr [][]key, i int, j int) (bool, bool) {
	for k := 0; k < len(arr[i - 1][j].p); k++ {
		for l := 0; l < len(arr[i][j - 1].p); l++ {
			if arr[i - 1][j].p[k].x == arr[i][j - 1].p[l].x && arr[i - 1][j].p[k].y == arr[i][j - 1].p[l].y {
				return true, true
			}
		}
	}
	return false, false
}
