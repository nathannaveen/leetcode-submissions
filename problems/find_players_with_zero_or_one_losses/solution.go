func findWinners(matches [][]int) [][]int {
    res := [][]int{ {}, {} }
    
    m := map[int] int {}
    
    for _, match := range matches {
        m[match[1]]++
        m[match[0]] += 0
    }
    
    for a, b := range m {
        if b == 0 {
            res[0] = append(res[0], a)
        } else if b == 1 {
            res[1] = append(res[1], a)
        }
    }
    
    sort.Ints(res[0])
    sort.Ints(res[1])
    
    return res
}