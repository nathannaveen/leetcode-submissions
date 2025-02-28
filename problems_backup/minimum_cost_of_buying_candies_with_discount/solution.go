func minimumCost(cost []int) int {
    sort.Slice(cost, func(i, j int) bool { return cost[i] > cost[j] })
    res := 0
    
    for i := 0; i < len(cost); i++ {
        if (i + 1) % 3 == 0 {
            continue
        }
        res += cost[i]
    }
    
    return res
}