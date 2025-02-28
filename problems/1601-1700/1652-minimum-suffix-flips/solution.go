func minFlips(target string) int {
    cur := 0
    res := 0
    
    for _, t := range target {
        if t == '0' && cur == 1 {
            cur = 0
            res++
        } else if t == '1' && cur == 0 {
            cur = 1
            res++
        }
    }
    
    return res
}