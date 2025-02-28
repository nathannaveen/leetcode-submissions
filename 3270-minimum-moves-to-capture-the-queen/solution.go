func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
    res := 100000
    if a == e {
        // fmt.Println(a, c, e, abs(a-c), abs(a - e))
        if a == c && abs(d - f) + abs(d - b) == abs(b - f) {
            res = 2
        } else {
            res = 1
        }
    } else if b == f {
        // fmt.Println(abs(a-c), abs(a - e))
        if b == d && abs(c - e) + abs(c - a) == abs(a - e) {
            res = 2
        } else {
            res = 1
        }
    } else {
        res = 2
    }
    
    dark := false
    if c % 2 != d % 2 {
        dark = true
    }
    
    qDark := false
    if e % 2 != f % 2 {
        qDark = true
    }
    
    if dark == qDark {
        if abs(e - c) == abs(f - d) { // one move
            if abs(c - a) == abs(d - b) && abs(a - e) + abs(a - c) == abs(c - e) && abs(b - f) + abs(b - d) == abs(d - f) {
                res = min(2, res)
            } else {
                res = 1
            }
        } else {
            res = min(res, 2)
        }
    }
    
    return res
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    
    return -a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}