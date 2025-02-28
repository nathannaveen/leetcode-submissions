func minimumDifference(nums []int, k int) int {
    min := 100000
    sort.Ints(nums)

    for i := k - 1; i < len(nums); i++ {
        min = int(math.Min(float64(min), float64(nums[i] - nums[i - k + 1]))) 
    }
    return min
}