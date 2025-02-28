func titleToNumber(columnTitle string) int {
    res := 0
    
    for _, letter := range columnTitle {
        res = res * 26 + int(letter - 64)
    }
    return res
}