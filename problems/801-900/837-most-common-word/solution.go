func mostCommonWord(paragraph string, banned []string) string {
    m := make(map[string] int)
    m2 := make(map[string] int)
    s := ""
    max := 0
    res := ""
    
    for _, ban := range banned {
        m[ban] = 1
    }
    
    for _, i := range paragraph {
        if i == ' ' || i == ',' || i == '!' || i == '?' || i == '\'' || i == ';' || i == '.' {
            m2[s]++
            s = ""
        } else {
            s += strings.ToLower(string(i))
        }
    }
    m2[s]++
    
    for a, b := range m2 {
        if b > max && m[a] == 0 && a != "" {
            max = b
            res = a
        }
    }
    
    return res
}