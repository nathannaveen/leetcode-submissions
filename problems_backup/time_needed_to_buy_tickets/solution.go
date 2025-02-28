func timeRequiredToBuy(tickets []int, k int) int {
    val := float64(tickets[k])
    res := 0
    
    for i := 0; i < len(tickets); i++ {
        res += int(math.Min(float64(tickets[i]), val))
        
        if i == k {
            val--
        }
    }
    return res
}