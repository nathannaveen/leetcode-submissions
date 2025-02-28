func countWays(nums []int) int {
    sort.Ints(nums)
    res := 0
    
    if nums[0] >= 1 {
        res++
    }
    
    for i := 0; i < len(nums); i++ {
        if i + 1 > nums[i] && (i == len(nums) - 1 || nums[i + 1] > i + 1) {
            res++
        }
    }
    
    return res
}