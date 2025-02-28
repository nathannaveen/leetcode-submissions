func meetRequirement(n int, lights [][]int, requirement []int) int {
    arr := make([]int, n)
    sum := 0
    res := 0
    
    for _, light := range lights {
        arr[max(0, light[0] - light[1])]++
        
        if light[0] + light[1] < n - 1 {
            arr[light[0] + light[1] + 1]--
        }
    }
    
    for i := range arr {
        sum += arr[i]
        if sum >= requirement[i] {
            res++
        }
    }
    
    return res
}

func max(a, b int) int {
    if a > b { return a }
    return b
}
