func minimumLength(s string) int {
    freqChar := map[rune] int {}
    
    for _, l := range s {
        freqChar[l]++
    }
    
    res := 0
    
    for _, v := range freqChar {
        if v % 2 == 0 {
            res += 2
        } else {
            res += 1
        }
    }
    
    return res
}
