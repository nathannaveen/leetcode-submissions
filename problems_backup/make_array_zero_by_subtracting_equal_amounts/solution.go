func minimumOperations(nums []int) int {
    m := map[int] bool {}
    
    for _, n := range nums {
        if n != 0 {
            m[n] = true
        }
    }
    
    return len(m)
}