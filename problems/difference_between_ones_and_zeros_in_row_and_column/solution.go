func onesMinusZeros(grid [][]int) [][]int {
    rows := make([]int, len(grid))
    cols := make([]int, len(grid[0]))

    res := make([][]int, len(grid))

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            val := grid[i][j]
            if grid[i][j] == 0 {
                val = -1
            }
            rows[i] += val
            cols[j] += val
        }
    }

    for i := 0; i < len(grid); i++ {
        res[i] = make([]int, len(grid[0]))
        for j := 0; j < len(grid[0]); j++ {
            res[i][j] = rows[i] + cols[j]
        }
    }

    return res
}