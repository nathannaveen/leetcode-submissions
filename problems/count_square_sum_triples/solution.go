func countTriples(n int) int {
    res := 0
    
    for a := 1; a < n; a++ {
        for b := 1; b < n; b++ {
            c := math.Sqrt(float64(a * a + b * b))
            intC := int(c)
            
            if c == float64(intC) && intC <= n {
                res++
            }
        }
    }
    
    return res
}