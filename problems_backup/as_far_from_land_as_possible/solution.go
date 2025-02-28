type pos struct {
    i int
    j int
}

func maxDistance(grid [][]int) int {
    n := len(grid)
    queue := []pos{}
    max := 1
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                queue = append(queue, pos{ i, j })
            }
        }
    }
    
    if len(queue) == 0 || len(queue) == n * n { return -1 }
    
    addToQueue := func(i, j, val int) {
        if i >= 0 && j >= 0 && i < n && j < n && grid[i][j] == 0 {
            grid[i][j] = val
            max = val
            queue = append(queue, pos{ i, j })
        }
    }
    
    for len(queue) != 0 {
        pop := queue[0]
        queue = queue[1:]
        i, j := pop.i, pop.j
        val := grid[i][j]
        
        addToQueue(i + 1, j, val + 1)
        addToQueue(i - 1, j, val + 1)
        addToQueue(i, j + 1, val + 1)
        addToQueue(i, j - 1, val + 1)
    }
       
    return max - 1
}