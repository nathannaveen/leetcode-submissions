func countNegatives(grid [][]int) int {
    row, col := 0, len(grid[0]) - 1
    res := 0

    for row < len(grid) && col >= 0 {
        if grid[row][col] >= 0 {
            row++
        } else {
            res += len(grid) - row
            col--
        }
    }
    
    return res
}