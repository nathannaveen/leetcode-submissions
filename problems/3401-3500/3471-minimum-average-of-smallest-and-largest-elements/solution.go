func minimumAverage(nums []int) float64 {
    minA := float64(10000)
    
    sort.Ints(nums)

    for i := 0; i < len(nums) / 2; i++ {
        x := float64(nums[i] + nums[len(nums) - 1 - i]) / 2
        if x < minA {
            minA = x
        }
    }

    return minA
}
