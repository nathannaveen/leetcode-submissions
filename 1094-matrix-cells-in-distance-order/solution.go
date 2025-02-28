func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
    type key struct {
        i, j int
    }
    
    res := [][]int{}
    
    queue := []key{ {rCenter, cCenter} }
    
    visited := map[key] bool { key{rCenter, cCenter} : true }
    
    appendToQueue := func(i, j int) {
        if i < 0 || j < 0 || i >= rows || j >= cols || visited[key{i, j}] {
            return
        }
        queue = append(queue, key{i, j})
        visited[key{i, j}] = true
    }
    
    for len(queue) > 0 {
        pop := queue[0]
        queue = queue[1:]
        
        i, j := pop.i, pop.j
        
        res = append(res, []int{i, j})

        appendToQueue(i + 1, j)
        appendToQueue(i - 1, j)
        appendToQueue(i, j + 1)
        appendToQueue(i, j - 1)
    }
    
    return res
}