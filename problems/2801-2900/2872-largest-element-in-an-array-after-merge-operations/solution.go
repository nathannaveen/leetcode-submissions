func maxArrayValue(nums []int) int64 {
    res := int64(0)

    for i := len(nums) - 1; i >= 1; i-- {
        if nums[i] >= nums[i - 1] {
            nums[i - 1] += nums[i]
            if int64(nums[i - 1]) > res {
                res =  int64(nums[i - 1])
            }
        }
    }

    if int64(nums[0]) > res {
        res = int64(nums[0])
    }

    return res
}