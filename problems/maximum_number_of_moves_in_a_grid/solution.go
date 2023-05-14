type key struct {
    i, j int
}

var dp map[key] int

func maxMoves(grid [][]int) int {
    res := 0
    dp = map[key] int {}
    
    for i := 0; i < len(grid); i++ {
        x := dfs(grid, i, 0)
        if x > res {
            res = x
        }
    }
    
    return res - 1
}

func dfs(grid [][]int, i, j int) int {
    // return max
    
    max := 0
    
    if val, ok := dp[key{i, j}]; ok {
        return val
    }
    
    for _, d := range [][]int{ {-1, 1}, {0, 1}, {1, 1} } {
        if inRange(grid, i + d[0], j + d[1]) && grid[i + d[0]][j + d[1]] > grid[i][j] {
            x := dfs(grid, i + d[0], j + d[1])
            if x > max {
                max = x
            }
        }
    }
    
    dp[key{i, j}] = 1 + max
    
    return 1 + max
}

func inRange(grid [][]int, i, j int) bool {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
        return false
    }
    return true
}