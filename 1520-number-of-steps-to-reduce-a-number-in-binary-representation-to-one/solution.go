func numSteps(s string) int {
    counter := 0
    one := 0
    
    for i := len(s) - 1; i >= 1; i-- {
        if one + int(s[i]) == 49 {
            counter++
            one = 1
        }
    }
    
    return len(s) - 1 + counter + one
}