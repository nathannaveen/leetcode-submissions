func strongPasswordCheckerII(password string) bool {
    lower, upper, digit, special := false, false, false, false
    
    for i := 0; i < len(password); i++ {
        l := password[i]
        if l < 123 && l > 96 {
            lower = true
        } else if l < 91 && l > 64 {
            upper = true
        } else if l < 58 && l > 47 {
            digit = true
        } else if strings.Contains("!@#$%^&*()-+", string(l)) {
            special = true
        }
        
        if i >= 1 && password[i - 1] == password[i] {
            return false
        }
    }
    
    return len(password) >= 8 && lower && upper && digit && special
}