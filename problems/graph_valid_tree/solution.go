func validTree(n int, edges [][]int) bool {
    if len(edges) != n - 1 { // checking for a loop
        return false
    }
    
    if n == 1 {
        return true
    }
    
    counter := 0
    queue := []int{ edges[0][0] }
    m := make(map[int] []int)
    contains := make(map[int] int)

    for _, i := range edges {
        m[i[0]] = append(m[i[0]], i[1])
        m[i[1]] = append(m[i[1]], i[0])
    }

    for len(queue) != 0 {
        pop := queue[0]
        queue = queue[1:]

        if contains[pop] != 1 {
            counter++
            queue = append(queue, m[pop]...)
            contains[pop] = 1
        }
    }

    return counter == n
}