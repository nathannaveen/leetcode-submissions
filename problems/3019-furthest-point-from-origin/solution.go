func furthestDistanceFromOrigin(moves string) int {
    l, r, x := 0, 0, 0
    
    for _, m := range moves {
        if m == 'L' {
            l++
        } else if m == 'R' {
            r++
        } else {
            x++
        }
    }
    
    if l > r {
        return l + x - r
    } else {
        return r + x - l
    }
}