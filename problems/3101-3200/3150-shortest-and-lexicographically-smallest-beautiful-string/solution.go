func shortestBeautifulSubstring(s string, k int) string {
    cnt := 0
    start := 0
    end := 0
    
    res := ""
    done := false
    
    for i := 0; i < len(s); i++ {
        end = i
        if s[i] == '1' {
            cnt++
            if cnt == k {
                done = true
                res = s[:i + 1]
                break
            }
        }
    }
    
    if done {
        for s[start] != '1' {
            start++
        }
        res = s[start :end + 1]
    }
    
    for i := end + 1; i < len(s); i++ {
        if s[i] == '1' {
            start++
            for s[start] != '1' {
                start++
            }
            
            if (done && ((s[start : i + 1] < res && len(s[start : i + 1]) <= len(res)) || len(s[start : i + 1]) < len(res))) || !done {
                res = s[start : i + 1]
            }
            
            done = true
        }
    }
    
    return res
}
