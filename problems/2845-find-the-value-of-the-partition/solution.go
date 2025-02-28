func findValueOfPartition(nums []int) int {
    sort.Ints(nums)
    res := 1000000000000
    
    for i := 1; i < len(nums); i++ {
        x := nums[i] - nums[i - 1]
        if x < res {
            res = x
        }
    }
    
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}