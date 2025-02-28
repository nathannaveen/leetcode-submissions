func minimumPossibleSum(n int, target int) int64 {
    i := 0
    x := 1
    m := map[int] bool {}
    sum := int64(0)
    
    for i < n {
        if !m[target - x] {
            i++
            m[x] = true
            sum += int64(x)
        }
        x++
    }
    
    return sum
}