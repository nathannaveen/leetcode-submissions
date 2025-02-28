func sumOfSquares(nums []int) int {
    res := 0
    
    for i := 1; i <= len(nums); i++ {
        if len(nums) % i == 0 {
            res += nums[i - 1] * nums[i - 1]
        }
    }
    
    return res
}