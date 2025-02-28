func findRepeatedDnaSequences(s string) []string {
    m := map[string] int {}
    res := []string{}
    
    for i := 0; i < len(s) - 9; i++ {
        x := s[i : i + 10]
        if m[x] == 1 {
            res = append(res, x)
        }
        m[x]++
    }

    return res
}