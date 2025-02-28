func numberOfPoints(nums [][]int) int {
    m := map[int] bool {}
    
    for _, n := range nums {
        for i := n[0]; i <= n[1]; i++ {
            m[i] = true
        }
    }
    
    return len(m)
}