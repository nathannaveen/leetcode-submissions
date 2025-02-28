func printVertically(s string) []string {
    arr := strings.Split(s, " ")
    max := 0
    
    for _, i := range arr {
        if len(i) > max { max = len(i) }
    }
    
    res := make([]string, max)
    
    for _, s := range arr {
        counter := 0
        for i := range res {
            if counter == len(s) {
                res[i] += " "
            } else {
                res[i] += string(s[counter])
                counter++
            }
        }
    }
    
    for i := range res {
        res[i] = strings.TrimRight(res[i], "\t \n")
    }
    
    return res
}