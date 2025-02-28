func numberOfPairs(nums []int) []int {
    m := map[int] bool {}
    res := 0
    
    for _, n := range nums {
        if m[n] {
            res++
            delete(m, n)
        } else {
            m[n] = true
        }
    }
    
    return []int{res, len(nums) - 2*res}
}