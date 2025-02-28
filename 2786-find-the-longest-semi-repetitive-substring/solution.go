func longestSemiRepetitiveSubstring(s string) int {
    max := 0
    prev := -1
    cnt := 1
    
    for i := 1; i < len(s); i++ {
        if s[i] == s[i - 1] {
            if prev != -1 {
                cnt = i - prev
            }
            prev = i
        }
        cnt++
        if cnt > max {
            max = cnt
        }
    }
    
    if cnt > max {
        max = cnt
    }
    
    return max
}