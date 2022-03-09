func minimalKSum(nums []int, k int) int64 {
    sort.Ints(nums)
    val := int64(1)
    res := int64(0)
    i := 0
    
    for j := 0; j < k; j++ {
        res += val
        
        if i < len(nums) && int64(nums[i]) <= val {
            res -= val
            val = int64(nums[i])
            j--
            i++
        }
        val++
    }
    
    return res
}