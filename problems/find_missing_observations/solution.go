func missingRolls(rolls []int, mean int, n int) []int {
    res := make([]int, n)
    sum := 0
    
    for _, roll := range rolls { 
        // finding sum
        sum += roll
    }
    
    missing := (len(rolls) + n) * mean - sum
    
    if missing > n * 6 || missing < n {
        // should return empty array
        return []int{}
    }
    
    for i := 0; i < n; i++ {
        // distribute the missing
        res[i] = int(math.Min(float64(6), float64(missing - (n - 1 - i))))
        missing -= res[i]
    }
    
    return res
}