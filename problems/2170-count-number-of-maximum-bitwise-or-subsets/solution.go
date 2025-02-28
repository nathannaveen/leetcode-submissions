func countMaxOrSubsets(nums []int) int {
    total := 0
    
    for _, n := range nums {
        total |= n
    }
    
    return helper(total, 0, 0, nums)
}

func helper(total, start, val int, nums []int) int {
    res := 0
    if val == total {
        res = 1
    }
    
    for i := start + 1; i <= len(nums); i++ {
        res += helper(total, i, val | nums[i - 1], nums)
    }
    
    return res
}