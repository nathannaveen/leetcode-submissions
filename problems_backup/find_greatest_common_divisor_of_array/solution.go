func findGCD(nums []int) int {
    min, max := 1001, 0
    
    for _, n := range nums {
        if n < min { min = n }
        if n > max { max = n }
    }
    
    for min != 0 {
        oldMin := min
        min = max % min
        max = oldMin
    }
    return max
}