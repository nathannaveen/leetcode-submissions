func distinctNames(ideas []string) int64 {
    m := map[int] map[string] int {} // whether m[prefix] can't use suffix
    
    for _, idea := range ideas {
        if _, ok := m[int(idea[0] - 'a')]; !ok {
            m[int(idea[0] - 'a')] = map[string] int {}
        }
        
        m[int(idea[0] - 'a')][idea[1:]] = 1
    }
    
    res := int64(0)
    
    for i := 0; i < 26; i++ {
        for j := i + 1; j < 26; j++ {
            cnt := 0
            
            for key := range m[i] {
                cnt += m[j][key]
            }
            
            res += int64(2) * int64((len(m[i]) - cnt) * (len(m[j]) - cnt))
        }
    }
    
    return res
}