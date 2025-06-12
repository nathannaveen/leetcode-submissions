type key struct {
    i, j, xor int
}

func countPathsWithXorValue(grid [][]int, k int) int {
    return helper(grid, k, 0, 0, grid[0][0], map[key] int {}) 
}

func helper(grid [][]int, k, i, j, val int, dp map[key] int) int {
    if i == len(grid) - 1 && j == len(grid[0]) - 1 && val == k {
        return 1
    }

    if val, ok := dp[key{i, j, val}]; ok {
        return val
    }

    res := 0

    if inBounds(grid, i + 1, j) {
        res += helper(grid, k, i + 1, j, val ^ grid[i + 1][j], dp)
    }
    if inBounds(grid, i, j + 1) {
        res += helper(grid, k, i, j + 1, val ^ grid[i][j + 1], dp)
    }

    res %= 1000000007

    dp[key{i, j, val}] = res

    return res
}

func inBounds(grid [][]int, i, j int) bool {
    return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}
