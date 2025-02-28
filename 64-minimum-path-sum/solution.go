type pos struct {
    i, j int
}

var dp = map[pos] int {}

func minPathSum(grid [][]int) int {
    dp = map[pos] int {}
    return helper(grid, 0, 0)
}

func helper(grid [][]int, i, j int) int {
    min := 8000000

    if i == len(grid) - 1 && j == len(grid[0]) - 1 {
        return grid[i][j]
    }

    if val, ok := dp[pos{i, j}]; ok {
        return val
    }

    for _, d := range []pos{ {0, 1}, {1, 0} } {
        I, J := i + d.i, j + d.j

        if I >= 0 && I < len(grid) && J >= 0 && J < len(grid[0]) {
            x := helper(grid, I, J)

            if x < min {
                min = x
            }
        }
    }

    dp[pos{i, j}] = grid[i][j] + min

    return grid[i][j] + min
}