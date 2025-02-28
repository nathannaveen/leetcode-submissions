func countAsterisks(s string) int {
    arr := strings.Split(s, "|")
    res := 0
    
    for i := 0; i < len(arr); i += 2 {
        res += strings.Count(arr[i], "*")
    }
    
    return res
}