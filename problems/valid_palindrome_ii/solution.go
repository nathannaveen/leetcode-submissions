func validPalindrome(s string) bool {
    i, j := 0, len(s) - 1
    
    for i < j {
        if s[i] != s[j] {
            return helper(i + 1, j, s) || helper(i, j - 1, s)
        }
        
        i++
        j--
    }
    
    return true
}

func helper(i, j int, s string) bool {
    for i < j {
        if s[i] != s[j] {
            return false
        }
        
        i++
        j--
    }
    
    return true
}