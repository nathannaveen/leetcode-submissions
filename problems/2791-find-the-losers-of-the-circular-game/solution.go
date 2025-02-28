func circularGameLosers(n int, k int) []int {
    m := map[int] bool { 1 : true }
    
    z := 1
    
    i := 0
    for {
        i++
        z += i * k
        x := z % n
        if x == 0 {
            x = n
        }
        // fmt.Println(z, x)
        if m[x] {
            break
        }
        m[x] = true
    }
    
    res := []int{}
    for i := 1; i <= n; i++ {
        if !m[i] {
            res = append(res, i)
        }
    }
    
    return res
}