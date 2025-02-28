func findClosestNumber(nums []int) int {
    res := 100001
    actual := 100001
    
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })
    
    for i := range nums {
        val := int(math.Abs(float64(nums[i])))
        if val < res {
            res = val
            actual = nums[i]
        }
    }
    
    return actual
}