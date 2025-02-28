type key struct {
    grid [][]int
    y int
    x int
    sum int
}

func getMaximumGold(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    stack := []key{}
    max := 0
    
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] != 0 {
                duplicate := make([][]int, n)
                for k := range grid {
                    duplicate[k] = make([]int, m)
                    copy(duplicate[k], grid[k])
                }
                
                add := duplicate[i][j]
                duplicate[i][j] = 0

                stack = append(stack, key{ duplicate, i, j, add })
            }
        }
    }
    
    mineGold := func(y, x, sum int, grid2 [][]int) {
        if x >= 0 && y >= 0 && x < m && y < n && grid2[y][x] != 0 {
            duplicate := make([][]int, n)
            for i := range grid2 {
                duplicate[i] = make([]int, m)
                copy(duplicate[i], grid2[i])
            }
            add := duplicate[y][x]
            duplicate[y][x] = 0
            stack = append(stack, key{ duplicate, y, x, sum + add })
        }
    }
    
    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        if pop.sum > max {
            max = pop.sum
        }
        
        mineGold(pop.y + 1, pop.x, pop.sum, pop.grid)
        mineGold(pop.y - 1, pop.x, pop.sum, pop.grid)
        mineGold(pop.y, pop.x + 1, pop.sum, pop.grid)
        mineGold(pop.y, pop.x - 1, pop.sum, pop.grid)
        
    }
    
    return max
}


