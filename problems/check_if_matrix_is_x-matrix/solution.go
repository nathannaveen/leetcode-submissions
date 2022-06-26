func checkXMatrix(grid [][]int) bool {
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if i + j == len(grid) - 1 || i == j {
                if grid[i][j] == 0 {
                    return false
                }
            } else if grid[i][j] != 0 {
                return false
            }
        }
    }
    return true
}