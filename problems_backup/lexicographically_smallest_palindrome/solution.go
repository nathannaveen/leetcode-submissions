func makeSmallestPalindrome(s string) string {
    res := ""
    res1 := ""
    
    for i := 0; i < len(s) / 2; i++ {
        if i == len(s) - 1 - i {
            res += string(s[i])
        } else if s[i] < s[len(s) - 1 - i] {
            res += string(s[i])
            res1 = string(s[i]) + res1
        } else if s[i] > s[len(s) - 1 - i] {
            res += string(s[len(s)- 1 - i])
            res1 = string(s[len(s)- 1 - i]) + res1
        } else {
            res += string(s[i])
            res1 = string(s[i]) + res1
        }
    }
    
    if len(s) % 2 == 1 {
        res += string(s[len(s) / 2])
    }
    
    return res + res1
}