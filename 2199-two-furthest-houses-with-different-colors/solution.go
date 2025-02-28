func maxDistance(colors []int) int {
    n := len(colors) - 1
    frontmost, endmost := 0, n
    
    for frontmost = 0; frontmost <= n; frontmost++ {
        if colors[n] != colors[frontmost] { break } 
    }
    
    for endmost = n; endmost >= 0; endmost-- { 
        if colors[0] != colors[endmost] { break } 
    }
    
    return int(math.Max(float64(n - frontmost), float64(endmost)))
}