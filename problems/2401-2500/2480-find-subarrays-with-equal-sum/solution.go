func findSubarrays(nums []int) bool {
    m := map[int] bool {}
    tmp := false
    
    for i := 0; i < len(nums) - 1; i++ {
        val := nums[i] + nums[i + 1]
        if tmp, m[val] = m[val], true; tmp {
            return true
        }
    }
    
    return false
}