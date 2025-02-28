func maximumDifference(nums []int) int {
    min := nums[0]
    max := -1
    
    for i := 0; i < len(nums); i++ {
        if nums[i] > min && nums[i] - min > max {
            max = nums[i] - min
        } else if nums[i] < min {
            min = nums[i]
        }
    }
    
    return max
}