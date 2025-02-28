func findMissingAndRepeatedValues(grid [][]int) []int {
    n := len(grid)
    
    res := []int{0, 0}
    x := 0
    
    m := map[int] bool {}
    cnt := 1
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            x += cnt
            if m[grid[i][j]] {
                res[0] = grid[i][j]
            } else {
                x -= grid[i][j]
            }
            m[grid[i][j]] = true
            cnt++
        }
    }
    
    res[1] = x
    
    return res
}