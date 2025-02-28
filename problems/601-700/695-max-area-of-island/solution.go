var doneAlready = make(map[[2]int] bool)

func maxAreaOfIsland(grid [][]int) int {
    res := 0
    doneAlready = make(map[[2]int] bool)
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            res = int(math.Max(float64(res), float64(helper(grid, i, j))))
        }
    }
    
    return res
}

func helper(grid [][]int, i, j int) int {
    
    /* 
    This part is for finding the area of the island, 
    but we can just check whether the area of the island 
    is greater than 0, if so we know there is an island.
    */
    temp := [2]int{i, j}
    
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || doneAlready[temp] || grid[i][j] == 0 {
        return 0
    }
    
    doneAlready[temp] = true
    
    return 1 + helper(grid, i - 1, j) + helper(grid, i + 1, j) + helper(grid, i, j - 1) + helper(grid, i, j + 1)
}