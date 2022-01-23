func countElements(nums []int) int {
    max, min := nums[0], nums[0]
    res := 0
    
    for _, n := range nums {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
    
    for _, n := range nums {
        if n < max && n > min {
            res++
        }
    }
    return res
}