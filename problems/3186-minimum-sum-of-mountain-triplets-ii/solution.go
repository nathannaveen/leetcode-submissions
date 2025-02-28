func minimumSum(nums []int) int {
    a, b := make([]int, len(nums)), make([]int, len(nums))
    aMin, bMin := math.MaxInt, math.MaxInt
    
    res := math.MaxInt
    
    for i := 0; i < len(nums); i++ {
        if nums[i] < aMin {
            aMin = nums[i]
        }
        a[i] = aMin
    }
    
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] < bMin {
            bMin = nums[i]
        }
        b[i] = bMin
    }
    
    for i := 1; i < len(nums) - 1; i++ {
        if a[i - 1] < nums[i] && b[i + 1] < nums[i] {
            x := a[i - 1] + nums[i] + b[i + 1]
            if x < res {
                res = x
            }
        }
    }
    
    if res == math.MaxInt {
        return -1
    }
    
    return res
}