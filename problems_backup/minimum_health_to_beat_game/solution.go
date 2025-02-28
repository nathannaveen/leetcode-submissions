func minimumHealth(damage []int, armor int) int64 {
    max := int64(0)
    res := int64(1)
    
    for _, d := range damage {
        res += int64(d)
        
        if int64(d) > max {
            max = int64(d)
        }
    }
    
    return res - int64(math.Min(float64(max), float64(armor)))
}