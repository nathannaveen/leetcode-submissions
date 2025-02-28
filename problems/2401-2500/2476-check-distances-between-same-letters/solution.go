func checkDistances(s string, distance []int) bool {
    m := map[int] int {}
    
    for i := 0; i < len(s); i++ {
        val := int(s[i] - 'a')
        
        if _, ok := m[val]; !ok {
            m[val] = i
        } else if distance[val] != i - m[val] - 1 {
            return false
        }
    }
    
    return true
}