func findMiddleIndex(nums []int) int {
    sum := 0
    left := 0
    
    for _, num := range nums { sum += num }
    
    for i := range nums {
        if sum - nums[i] - left == left {
            return i
        }
        
        left += nums[i]
    }
    
    return -1
}