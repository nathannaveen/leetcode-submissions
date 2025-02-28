func convertToBase7(num int) string {
    neg := false
    res := ""
    
    if num < 0 {
        neg = true
        num *= -1
    }
    
    for num > 0 {
        res = strconv.Itoa(num % 7) + res
        num /= 7
    }
    
    if res == "" {
        return "0"
    }
    
    if neg {
        res = "-" + res
    }
    
    return res
}