func trap(height []int) int {
    res := 0
    lToR := make([]int, len(height))
    rToL := make([]int, len(height))
    max := height[0]
    
    for i := 0; i < len(height); i++ {
        if height[i] > max { max = height[i] }
        lToR[i] = max
    }
    
    max = height[len(height) - 1]
    
    for i := len(height) - 1; i >= 0; i-- {
        if height[i] > max { max = height[i] }
        rToL[i] = max
    }
    
    for i := 0; i < len(height); i++ {        
        res += int(math.Min(float64(rToL[i]), float64(lToR[i]))) - height[i]
    }
    
    return res
}