func maxNumberOfBalloons(text string) int {
    m := map[rune] int {}
    min := 1429
    
    for _, l := range text {
        m[l]++
    }
    
    for _, l := range "ban" {
        min = minimum(min, m[l])
    }
    for _, l := range "ol" {
        min = minimum(min, m[l] / 2)
    }
    
    return min
}

func minimum(a, b int) int {
    if a < b { 
        return a 
    }
    return b
}