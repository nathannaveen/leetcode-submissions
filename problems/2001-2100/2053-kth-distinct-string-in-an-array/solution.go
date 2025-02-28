func kthDistinct(arr []string, k int) string {
    m := make(map[string] int)
    counter := 0
    
    for _, a := range arr { m[a]++ }
    
    for _, a := range arr {
        if m[a] == 1 { counter++ }
        if counter == k { return a }
    }
    
    return ""
}