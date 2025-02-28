func countPoints(rings string) int {
    rods := make([]int, 10)
    res := 0
    
    for i := 0; i < len(rings); i += 2 {
        pos, _ := strconv.Atoi(string(rings[i + 1]))
        
        if rods[pos] == 0 { rods[pos] = 1 }
        
        if rings[i] == 'R' {
            rods[pos] *= 2
        } else if rings[i] == 'G' {
            rods[pos] *= 3
        } else if rings[i] == 'B' {
            rods[pos] *= 5
        }
    }
    
    for _, rod := range rods {
        if rod % 2 == 0 && rod % 3 == 0 && rod % 5 == 0 && rod != 0 { res++ }
    }
    
    return res
}