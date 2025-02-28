func removeOccurrences(s string, part string) string {
    return a(s, part, 0)
}

func a(s, part string, i int) string {
    if i > -1 && i <= len(s) - len(part) && s[i : i + len(part)] == part {
        s = a(s[0 : i] + s[i + len(part) : len(s)], part, i - len(part))
    } else if i > len(s) - len(part) {
        return s
    }
    s = a(s, part, i + 1)
    
    return s
}