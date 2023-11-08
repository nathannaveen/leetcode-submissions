func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
    if sx == fx && sy == fy {
        return t != 1
    }
    xDiff, yDiff := abs(fx - sx), abs(fy - sy)
    min := int(math.Min(float64(xDiff), float64(yDiff)))
    z := xDiff - min + yDiff
    
    return t >= z
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}