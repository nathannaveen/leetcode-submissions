func isPathCrossing(path string) bool {
    x, y := 0, 0
    m := map[[2]int] bool { [2]int{ 0, 0 } : true }
    
    for _, p := range path {
        switch p {
            case 'N':
                y++
            case 'S':
                y--
            case 'E':
                x++
            case 'W':
                x--
        }
        
        pos := [2]int{ x, y }
        
        if m[pos] {
            return true
        }
        m[pos] = true
    }
    
    return false
}