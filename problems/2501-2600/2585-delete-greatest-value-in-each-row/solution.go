func deleteGreatestValue(grid [][]int) int {
    res := 0

    for _, row := range grid {
        sort.Slice(row, func(i, j int) bool { return row[i] > row[j] })
    }

    for len(grid[0]) > 0 {
        max := 0

        for i := 0; i < len(grid); i++ {
            if grid[i][0] > max {
                max = grid[i][0]
            }
            grid[i] = grid[i][1:]
        }

        res += max
    }

    return res
}