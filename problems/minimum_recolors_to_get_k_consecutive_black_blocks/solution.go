func minimumRecolors(blocks string, k int) int {
    c := 0
    res := 0
    
    for i := 0; i < k; i++ {
        if blocks[i] == 'B' {
            c++
        }
    }
    
    res = max(c, res)
    
    for i := k; i < len(blocks); i++ {
        if blocks[i - k] == 'B' {
            c--
        }
        if blocks[i] == 'B' {
            c++
        }
        
        res = max(c, res)
    }
    
    return k - res
}

func max(a, b int) int {
    if a < b {
        return b
    }
    
    return a
}