func camelMatch(queries []string, pattern string) []bool {
    res := []bool{}
    
    for _, q := range queries {
        c := 0
        canDo := true
        
        for i := 0; i < len(q); i++ {
            if c != len(pattern) && pattern[c] == q[i] {
                c++
            } else if q[i] > 64 && q[i] < 91 {
                canDo = false
                break
            }
        }
        
        if c != len(pattern) {
            canDo = false
        }
        
        res = append(res, canDo)
    }
    
    return res
}