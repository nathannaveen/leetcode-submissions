func minimumSemesters(n int, relations [][]int) int {
    numOfPre := make([]int, n)
    m := make(map[int] []int)
    
    for _, relation := range relations {
        prerequisite := relation[0]
        course := relation[1]
        
        m[prerequisite] = append(m[prerequisite], course)
        numOfPre[course - 1]++
    }
    
    queue := []int{}
    res := 0
    counter := 0
    
    for i := 0; i < n; i++ {
        if numOfPre[i] == 0 { queue = append(queue, i + 1) }
    }
    
    for len(queue) != 0 {
        res++
        n := len(queue)
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            counter++
            
            for _, temp := range m[pop] {
                numOfPre[temp - 1]--
                if numOfPre[temp - 1] == 0 {
                    queue = append(queue, temp)
                }
            }
        }
    }
    
    if counter == n { return res }
    return -1
}