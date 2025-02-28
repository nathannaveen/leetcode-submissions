func restoreArray(adjacentPairs [][]int) []int {
    m := map[int] []int {}
    used := map[int] bool {}
    c := 100001
    res := []int{}
    
    for _, a := range adjacentPairs {
        m[a[0]] = append(m[a[0]], a[1])
        m[a[1]] = append(m[a[1]], a[0])
    }
    
    for a, b := range m {
        if len(b) == 1 {
            c = a
            break
        }
    }
    
    res = append(res, c)
    used[c] = true
    
    for c != 100001 {
        l := 100001
        
        for _, b := range m[c] {
            if !used[b] {
                l = b
                break
            }
        }
        
        used[c] = true
        
        c = l
        
        res = append(res, c)
    }
    
    return res[:len(res) - 1]
}