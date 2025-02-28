func digitCount(num string) bool {
    m := map[int] int {}
    
    for _, digit := range num {
        m[int(digit - '0')]++
    }
    
    for i := 0; i < len(num); i++ {
        n := strconv.Itoa(m[i])
        
        if n != string(num[i]) {
            return false
        }
    }
    
    return true
}