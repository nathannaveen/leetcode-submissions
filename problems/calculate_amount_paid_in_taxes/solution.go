func calculateTax(brackets [][]int, income int) float64 {
    res := float64(0)
    prevBracket := 0
    
    for i := 0; i < len(brackets); i++ {
        tax := float64(brackets[i][1]) / float64(100)
        res += float64(int(math.Min(float64(brackets[i][0]), float64(income))) - prevBracket) * tax
        
        prevBracket = brackets[i][0]
        
        if income <= brackets[i][0] {
            break
        }
    }
    
    return res
}