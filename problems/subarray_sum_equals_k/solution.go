func subarraySum(nums []int, k int) int {
    m := map[int] int{ 0 : 1 }
    sum := 0
    res := 0
    
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if m[sum - k] > 0 {
            res += m[sum - k]
        }
        m[sum]++
    }
    
    return res
}