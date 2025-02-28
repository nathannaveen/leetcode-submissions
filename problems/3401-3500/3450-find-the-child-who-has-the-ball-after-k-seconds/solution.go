func numberOfChild(n int, k int) int {
    c := 0
    add := 0
    for i := 0; i < k; i++ {
        if c == 0 {
            add = 1
        } else if c == n - 1 {
            add = -1
        }
        c += add
    }
    
    return c
}
