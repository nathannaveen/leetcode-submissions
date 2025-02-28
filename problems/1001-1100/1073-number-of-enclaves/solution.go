var doneAlready = make(map[[2]int] bool)

func numEnclaves(grid [][]int) int {
    res := 0
    doneAlready = make(map[[2]int] bool)
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            temp := helper(grid, i, j)
            if temp > 0 {
                res += temp
            }
        }
    }
    
    return res
}

func helper(grid [][]int, i, j int) int {
    
    temp := [2]int{i, j}
    
    if doneAlready[temp] || grid[i][j] == 0 {
        return 0
    }
    
    if i <= 0 || i >= len(grid) - 1 || j <= 0 || j >= len(grid[0]) - 1 {
        return -250001
    }
    
    doneAlready[temp] = true
    
    return 1 + helper(grid, i - 1, j) + helper(grid, i + 1, j) + helper(grid, i, j - 1) + helper(grid, i, j + 1)
}