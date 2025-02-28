func isStrictlyPalindromic(n int) bool {
    for i := 2; i <= n - 2; i++ {
        if !isPalindromic(strconv.FormatInt(int64(n), i)) {
            return false
        }
    }
    
    return true
}

func isPalindromic(s string) bool {
    return s == reverse(s)
}

func reverse(s string) (res string) {
    for _, l := range s {
        res = string(l) + res
    }
    
    return
}