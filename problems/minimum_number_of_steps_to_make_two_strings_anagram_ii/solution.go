func minSteps(s string, t string) int {
    letters := make([]int, 26)
    res := 0
    
    for _, letter := range s { letters[letter - 'a']++ }
    for _, letter := range t { letters[letter - 'a']-- }
    for _, n := range letters { res += int(math.Abs(float64(n))) }
    
    return res
}