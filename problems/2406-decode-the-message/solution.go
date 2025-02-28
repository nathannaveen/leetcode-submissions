func decodeMessage(key string, message string) string {
    found := map[string] bool {}
    
    arr := make([]string, 26)
    
    for _, l := range key {
        if !found[string(l)] && l != ' ' {
            arr[l - 'a'] = string(rune('a' + len(found)))
            found[string(l)] = true
        }
    }
    
    res := ""
    
    for _, l := range message {
        if l == ' ' {
            res += " "
        } else {
            res += arr[l - 'a']
        }
    }
    
    return res
}