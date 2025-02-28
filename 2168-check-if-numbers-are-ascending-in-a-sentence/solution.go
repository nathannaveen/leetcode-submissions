func areNumbersAscending(s string) bool {
    arr := strings.Split(s, " ")
    prev := -1
    
    for _, a := range arr {
        n, _ := strconv.Atoi(a)
        
        if prev < n { // check if it is a digit
            prev = n
            continue
        } else if n != 0 { 
            return false
        }
    }
    return true
}