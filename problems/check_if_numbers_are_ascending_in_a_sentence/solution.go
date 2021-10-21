func areNumbersAscending(s string) bool {
    arr := strings.Split(s, " ")
    prev := 0
    
    for _, a := range arr {
        n, _ := strconv.Atoi(a)
        
        if n != 0 && prev < n { // check if it is a digit
            prev = n
            continue
        } else if n != 0 { 
            return false
        }
    }
    return true
}