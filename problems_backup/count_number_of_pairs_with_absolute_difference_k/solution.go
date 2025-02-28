func countKDifference(nums []int, k int) int {
    res := 0
    m := make(map[int] int)
    
    for _, num := range nums {
        m[num]++
        res += m[num + k] + m[num - k]
    }
    
    return res
}