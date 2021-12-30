func validTree(n int, edges [][]int) bool {
    if len(edges) != n - 1 { // checking for a loop
        return false
    }
    
    if n == 1 {
        return true
    }
    
    counter := 0
    stack := []int{ edges[0][0] }
    m := make(map[int] []int)
    contains := make(map[int] int)

    for _, i := range edges {
        m[i[0]] = append(m[i[0]], i[1])
        m[i[1]] = append(m[i[1]], i[0])
    }

    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]

        if contains[pop] != 1 {
            counter++
            stack = append(stack, m[pop]...)
            contains[pop] = 1
        }
    }

    return counter == n
}