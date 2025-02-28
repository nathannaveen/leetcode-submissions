func removeKdigits(num string, k int) string {
    res := ""
    
    for i := 0; i < len(num); i++ {
        for len(res) > 0 && res[len(res) - 1] > num[i] && k > 0 {
            res = res[:len(res) - 1]
            k--
        }
        res += string(num[i])
    }
    
    for k > 0 {
        k--
        res = res[:len(res) - 1]
    }
    
    i := 0
    for i = 0; i < len(res); i++ {
        if res[i] != '0' { break }
    }
    
    res = res[i:]
    
    if len(res) != 0 {
        return res
    }
    
    return "0"
}