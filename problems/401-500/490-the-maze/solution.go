type point struct {
    i, j int
}

func hasPath(maze [][]int, start []int, destination []int) bool {
    stack := []point{ point{ start[0], start[1] } }
    visited := make(map[point] bool)
    n, m := len(maze), len(maze[0])
    
    helper := func(moveI, moveJ, i, j int) {
        for i < n && j < m && i >= 0 && j >= 0 && maze[i][j] != 1 {
            i += moveI
            j += moveJ
        }
        i -= moveI
        j -= moveJ
        
        if !visited[point{ i, j }] {
            stack = append(stack, point{ i, j })
            visited[point{ i, j }] = true
        }
    }
    
    for len(stack) != 0 {
        pop := stack[0]
        stack = stack[1:]
        
        if pop.i == destination[0] && pop.j == destination[1] {
            return true
        }
        
        helper(1, 0, pop.i, pop.j)
        helper(-1, 0, pop.i, pop.j)
        helper(0, 1, pop.i, pop.j)
        helper(0, -1, pop.i, pop.j)
        
    }
    
    return false
}