func minLength(s string) int {
    for i := 0; i < len(s) - 1; i++ {
        // fmt.Println(i)
        if i >= 0 {
            // fmt.Println(s[i : i + 2])
            if s[i : i + 2] == "AB" || s[i : i + 2] == "CD" {
                s = s[:i] + s[i + 2 :]
                i -= 2
                // fmt.Println("remove", s)
            }
        }
    }
    
    return len(s)
}