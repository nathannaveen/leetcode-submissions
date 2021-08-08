func addRungs(rungs []int, dist int) int {
    res := 0
    rungs = append(rungs[:0], append([]int{0}, rungs[0:]...)...)
    
    for i := 1; i < len(rungs); i++ {
        res += (rungs[i] - rungs[i - 1] - 1) / dist
    }
    return res
}