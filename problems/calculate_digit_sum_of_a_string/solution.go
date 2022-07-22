func digitSum(s string, k int) string {
    for len(s) > k {
        newS := ""
        
        for i := 0; i < len(s); i += k {
            group := s[i : int(math.Min(float64(i + k), float64(len(s))))]
            n := 0
            
            for _, digit := range group {
                num, _ := strconv.Atoi(string(digit))
                n += num
            }
            
            newS += strconv.Itoa(n)
        }
        
        s = newS
    }
    
    return s
}