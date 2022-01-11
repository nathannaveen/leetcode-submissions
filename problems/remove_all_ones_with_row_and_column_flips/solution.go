func removeOnes(grid [][]int) bool {
    n, m := len(grid), len(grid[0])
    invert := make([]int, len(grid[0]))
    
    for i := 0; i < m; i++ {
        invert[i] = -1 * (grid[0][i] - 1)
    }
    
    for i := 0; i < n; i++ {
        if reflect.DeepEqual(grid[i], grid[0]) || reflect.DeepEqual(grid[i], invert) {
            continue
        }
        return false
    }
    return true
}