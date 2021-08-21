func validTree(n int, edges [][]int) bool {
    counter := 0
    
    if len(edges) != n - 1 {
        return false
    }
    
    if len(edges) > 0 {
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
        
        if counter != n {
            return false
        }
    }
    return true
}