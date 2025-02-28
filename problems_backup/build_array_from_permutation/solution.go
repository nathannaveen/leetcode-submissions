func buildArray(nums []int) []int {
    res := make([]int, len(nums))
    
    for i := 0; i < len(nums); i++ {
        res[i] = nums[nums[i]]
    }
    
    return res
}