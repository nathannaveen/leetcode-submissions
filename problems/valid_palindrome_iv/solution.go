func makePalindrome(s string) bool {
    l, r := 0, len(s) - 1
    c := 0
    
    for l < r {
        if s[l] != s[r] {
            c++
        }
        l++
        r--
    }
    
    return c <= 2
}