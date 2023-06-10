func isFascinating(n int) bool {
    a, b := n * 2, n * 3
    m := map[int] bool {}
    
    for n > 0 {
        if n % 10 == 0 {
            return false
        } else if m[n % 10] {
            return false
        }
        m[n % 10] = true
        n /= 10
    }
    for a > 0 {
        if a % 10 == 0 {
            return false
        } else if m[a % 10] {
            return false
        }
        m[a % 10] = true
        a /= 10
    }
    for b > 0 {
        if b % 10 == 0 {
            return false
        } else if m[b % 10] {
            return false
        }
        m[b % 10] = true
        b /= 10
    }
    return len(m) == 9 
}