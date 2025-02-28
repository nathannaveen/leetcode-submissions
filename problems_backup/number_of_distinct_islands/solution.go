type key struct {
    i, j int
}

var visited = map[key] bool {}

func numDistinctIslands(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    visited = map[key] bool {}
    arr := []string{}
    res := 0
    
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 {
                temp := helper(i, j, grid, "")
                add := true
                for _, s := range arr {
                    if s == temp {
                        add = false
                        break
                    }
                }
                if add {
                    res++
                }
                arr = append(arr, temp)
            }
        }
    }
    
    return res
}

func helper(i, j int, grid [][]int, move string) string {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
        return "/"
    }
    if visited[key{ i, j }] {
        return ""
    }
    
    visited[key{ i, j }] = true
    grid[i][j] = 2
    
    res := move
    
    res += helper(i, j + 1, grid, "R")
    res += helper(i, j - 1, grid, "L")
    res += helper(i - 1, j, grid, "U")
    res += helper(i + 1, j, grid, "D")
    
    return res
}