func countTime(time string) int {
    h1, h2, m1, m2 := time[0], time[1], time[3], time[4]
    res := 1
    
    if m1 == '?' {
        res *= 6
    } 
    if m2 == '?' {
        res *= 10
    }
    
    if h1 == '?' && h2 == '?' {
        res *= 24
    } else if h1 == '?' {
        if int(h2) - 48 >= 4 {
            res *= 2
        } else {
            res *= 3
        }
    } else if h2 == '?' {
        if int(h1) - 48 == 2 {
            res *= 4
        } else {
            res *= 10
        }
    }
    
    return res
}