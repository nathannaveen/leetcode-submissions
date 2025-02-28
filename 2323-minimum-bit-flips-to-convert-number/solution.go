func minBitFlips(start int, goal int) int {
    res := 0
    
    for start > 0 || goal > 0 {
        if start & 1 != goal & 1 {
            res++
        }
        start >>= 1
        goal >>= 1
    }
    
    return res
}