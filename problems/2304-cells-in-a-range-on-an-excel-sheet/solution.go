func cellsInRange(s string) []string {
    res := []string{}
    c1, r1, c2, r2 := s[0], s[1], s[3], s[4] 
    // len(s) == 5 and '1' <= s[1] <= s[4] <= '9' so we can do the above
    
    for c := c1; c <= c2; c++ {
        for r := r1; r <= r2; r++ {
            res = append(res, string(c) + string(r))
        }
    }
    
    return res
}