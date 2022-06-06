func removeDigit(number string, digit byte) string {
    index := -1
    d, _ := strconv.Atoi(string(digit))
    
    for i := 0; i < len(number); i++ {
        n, _ := strconv.Atoi(string(number[i]))
        n2 := 0
        
        if i != len(number) - 1 {
            n2, _ = strconv.Atoi(string(number[i + 1]))
        }
        
        if n != d {
            continue
        }
        
        if n2 > n {
            return number[: i] + number[i + 1 :]
        }
        
        index = i
        
    }
    
    return number[: index] + number[index + 1 :]
}