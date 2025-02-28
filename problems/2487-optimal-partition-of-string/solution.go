func partitionString(s string) int {
    seen := map[byte] bool {}
    res := 0
    
    for i := 0; i < len(s); i++ {
        l := s[i]
        
        if seen[l] {
            res++
            seen = map[byte] bool { l : true }
        } else {
            seen[l] = true
        }
    }
    
    return res + 1
}