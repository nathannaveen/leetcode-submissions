type pos struct {
    i, j int
}

var canDo = true

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    n, m := len(grid1), len(grid1[0])
    res := 0
    visited := map[pos] bool {}

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if !visited[pos{i, j}] && grid2[i][j] != 0 {
                canDo = true
                findSubIsland(i, j, grid1, grid2, visited)
                if canDo {
                    res++
                }
            }
        }
    }

    return res
}

func findSubIsland(i, j int, grid, grid2 [][]int, visited map[pos] bool) {
    if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid2[i][j] == 0 || visited[pos{i, j}] {
        return
    }

    if grid[i][j] == 0 {
        canDo = false
    }

    visited[pos{i, j}] = true
    
    findSubIsland(i + 1, j, grid, grid2, visited)
    findSubIsland(i - 1, j, grid, grid2, visited)
    findSubIsland(i, j + 1, grid, grid2, visited)
    findSubIsland(i, j - 1, grid, grid2, visited)
}