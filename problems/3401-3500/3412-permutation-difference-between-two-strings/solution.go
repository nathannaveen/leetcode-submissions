func findPermutationDifference(s string, t string) int {
    m := map[rune] int {}
    res := 0
    
    for i, l := range s {
        m[l] = i
    }
    
    for i, l := range t {
        res += abs(m[l] - i)
    }
    
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
