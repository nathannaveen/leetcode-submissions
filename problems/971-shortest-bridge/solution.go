type point struct {
    i, j int
}

func shortestBridge(grid [][]int) int {
    queue := []point{}
    res := 0
    
    for i := 0; i < len(grid); i++ { // find first point
        shouldBreak := false
        for j := 0; j < len(grid); j++ { 
            if grid[i][j] == 1 {
                queue = dfs(i, j, grid) // dfs on that first point
                shouldBreak = true
                break
            }
        }
        if shouldBreak {
            break
        }
    }
    
    for len(queue) != 0 {
        n := len(queue)
        
        for k := 0; k < n; k++ {
            pop := queue[0]
            queue = queue[1:]
            i, j := pop.i, pop.j
            
            if i < 0 || j < 0 || i >= len(grid) || j >= len(grid) || grid[i][j] == -2 { 
                continue
            }
            if grid[i][j] == 1 {
                return res - 1
            }
            grid[i][j] = -2
            queue = append(queue, point{i + 1, j}, point{i - 1, j}, point{i, j + 1}, point{i, j - 1})
        }
        res++
    }
    return 0
}

func dfs(i, j int, grid [][]int) []point {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid) || grid[i][j] != 1 {
        return []point{}
    }
    res := []point{ point{ i, j } }
    grid[i][j] = -1
    
    res = append(res, dfs(i + 1, j, grid)...)
    res = append(res, dfs(i - 1, j, grid)...)
    res = append(res, dfs(i, j + 1, grid)...)
    res = append(res, dfs(i, j - 1, grid)...)
    
    return res
}