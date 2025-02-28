func minimumSum(n int, k int) int {
    m := map[int] bool {}
    
    i := 0
    s := 0
    
    for n > 0 {
        i++
        if m[k - i] {
            continue
        }
        n--
        s += i
        m[i] = true
    }
    
    return s
}