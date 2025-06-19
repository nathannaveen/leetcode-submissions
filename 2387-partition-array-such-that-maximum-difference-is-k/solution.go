func partitionArray(nums []int, k int) int {
    sort.Ints(nums)

    res := 0
    min := nums[0]

    for i := 0; i < len(nums); i++ {
        if nums[i] - min > k {
            res++
            min = nums[i]
        }
    }

    return res + 1
}
