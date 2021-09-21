type point struct {
    i, j int
}

var doneAlready = make(map[point] bool)

func numIslands(grid [][]byte) int {
    res := 0
    doneAlready = make(map[point] bool)
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if helper(grid, i, j) {
                res++
            }
        }
    }
    
    return res
}

func helper(grid [][]byte, i, j int) bool {
    
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || 
        doneAlready[point{ i, j }] || grid[i][j] == '0' {
        
        return false
    }
    
    doneAlready[point{ i, j }] = true
    
    helper(grid, i - 1, j)
    helper(grid, i + 1, j)
    helper(grid, i, j - 1)
    helper(grid, i, j + 1)
    
    return true
}