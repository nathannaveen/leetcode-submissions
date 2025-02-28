func longestContinuousSubstring(s string) int {
    c := 0
    prev := 0
    res := 0
    
    for _, l := range s {
        if int(l) == prev + 1 {
            c++
            prev = int(l)
        } else {
            res = int(math.Max(float64(res), float64(c)))
            prev = int(l)
            c = 1
        }
    }
    
    res = int(math.Max(float64(res), float64(c)))
    
    return res
}