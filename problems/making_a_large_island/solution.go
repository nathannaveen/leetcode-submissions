func largestIsland(grid [][]int) int {
    val := 2
    m := make(map[int] int)
    n := len(grid)
    res := 0
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                m[val] = dfs(i, j, val, grid)
                val++
            }
        }
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                visited := make(map[int] bool)
                sum := 1
                
                helper := func(i, j int) {
                    if i < 0 || j < 0 || i >= n || j >= n || visited[grid[i][j]] {
                        return
                    }
                    visited[grid[i][j]] = true
                    sum += m[grid[i][j]]
                }
                
                helper(i + 1, j)
                helper(i - 1, j)
                helper(i, j + 1)
                helper(i, j - 1)
                
                if sum > res {
                    res = sum
                }
            }
        }
    }
    
    for _, b := range m {
        if b > res {
            res = b
        }
    }
    
    return res
}

func dfs(i, j, val int, grid [][]int) int {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid) || grid[i][j] != 1 {
        return 0
    }
    grid[i][j] = val
    
    res := 1
    res += dfs(i + 1, j, val, grid)
    res += dfs(i - 1, j, val, grid)
    res += dfs(i, j + 1, val, grid)
    res += dfs(i, j - 1, val, grid)
    
    return res
}