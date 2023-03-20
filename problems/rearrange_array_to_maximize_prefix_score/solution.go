func maxScore(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })

    sum := 0
    res := 0

    for _, num := range nums {
        sum += num
        if sum <= 0 {
            return res
        }
        res++
    }

    return res
}