func findColumnWidth(grid [][]int) []int {
    res := make([]int, len(grid[0]))
    
    for j := 0; j < len(grid[0]); j++ {
        max := 0
        for i := 0; i < len(grid); i++ {
            s := strconv.Itoa(grid[i][j])
            if len(s) > max {
                max = len(s)
            }
        }
        res[j] = max
    }
    
    return res
}