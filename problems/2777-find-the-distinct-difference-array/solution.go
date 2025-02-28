func distinctDifferenceArray(nums []int) []int {
    res := make([]int, len(nums))
    
    for i := 0; i < len(nums); i++ {
        distinct := map[int] int {}
        a, b := 0, 0
        
        for j := 0; j <= i; j++ {
            distinct[nums[j]]++
        }
        a = len(distinct)
        distinct = map[int] int {}
        
        for j := i + 1; j < len(nums); j++ {
            distinct[nums[j]]++
        }
        b = len(distinct)
        
        res[i] = a - b
    }
    
    return res
}