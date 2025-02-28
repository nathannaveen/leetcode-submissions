func maxSum(grid [][]int) int {
    res := 0
    
    for i := 0; i < len(grid) - 2; i++ {
        for j := 0; j < len(grid[0]) - 2; j++ {
            c := grid[i][j] + grid[i + 2][j] + 
            grid[i][j + 1] + grid[i + 1][j + 1] + grid[i + 2][j + 1] + 
            grid[i][j + 2] + grid[i + 2][j + 2]
            
            if c > res {
                res = c
            }
        }
    }
    
    return res
}