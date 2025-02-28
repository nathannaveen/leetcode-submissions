func minimumSwaps(nums []int) int {
    min, max := 100001, 0
    minI, maxI := 0, 0
    
    for i, num := range nums {
        if num < min {
            min = num
            minI = i
        }
        if num >= max {
            max = num
            maxI = i
        }
    }
    
    remove := 0
    
    if minI > maxI {
        remove = 1
    }
    
    return minI + (len(nums) - 1 - maxI) - remove
}