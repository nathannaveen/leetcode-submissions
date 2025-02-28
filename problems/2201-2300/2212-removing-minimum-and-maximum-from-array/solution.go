func minimumDeletions(nums []int) int {
    minimum, maximum := 100000, -100000
    minI, maxI := 0, 0
    n := len(nums)
    
    for i, num := range nums {
        if num < minimum {
            minimum = num
            minI = i
        }
        
        if num > maximum {
            maximum = num
            maxI = i
        }
    }
    
    if minI > maxI {
        minI, maxI = maxI, minI
    }
    
    return min(minI + 1 + (n - maxI), min(n - minI, maxI + 1))
}

func min(a, b int) int {
    if a < b { return a }
    return b
}