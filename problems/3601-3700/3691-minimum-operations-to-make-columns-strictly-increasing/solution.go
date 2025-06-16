func minimumOperations(grid [][]int) int {
    cnt := 0
    for i := 1; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i - 1][j] >= grid[i][j] {
                x := grid[i - 1][j] - grid[i][j] + 1
                cnt += x
                grid[i][j] += x
            }
        }
    }

    return cnt
}
