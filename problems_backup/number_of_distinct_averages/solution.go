func distinctAverages(nums []int) int {
    sort.Ints(nums)
    m := map[float64] int {}
    
    for i := 0; i < len(nums) / 2; i++ {
        average := float64(nums[i] + nums[len(nums) - 1 - i]) / float64(2)
        m[average]++
    }
    return len(m)
}