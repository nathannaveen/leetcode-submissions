func countPairs(nums []int, k int) int {
    res := 0
    
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] && i * j % k == 0 {
                res++
            }
        }
    }
    
    return res
}