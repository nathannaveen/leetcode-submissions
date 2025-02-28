func myAtoi(s string) int {
    s = strings.TrimLeft(s, " ")
    res := 0
    neg := 1
    i := 0
    
    if len(s) == 0 {
        return 0
    }
    
    if s[0] == '-' {
        neg = -1
        i++
    } else if s[0] == '+' {
        i++
    }
    
    for i < len(s) {
        char := s[i]
        i++
        
        if res > 2147483647 {
            break
        }
        
        if int(char - '0') > -1 && int(char - '0') < 10 {
            res = res * 10 + int(char - '0')
            continue
        }
        break
    }
    
    res *= neg
    
    if res < -2147483648 {return -2147483648}
    if res > 2147483647 {return 2147483647}
    
    return res
}