func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
    res := int64(0)
    
    for i := 0; i <= total / cost2; i++ {
        res += int64(total - (cost2 * i)) / int64(cost1) + 1
    }
    
    return res
}