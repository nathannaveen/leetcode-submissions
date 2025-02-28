func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    res := 0
    
    for _, h := range hours {
        if h >= target {
            res++
        }
    }
    
    return res
}