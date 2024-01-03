func largestPerimeter(nums []int) int64 {
    sort.Ints(nums)

    res := int64(0)
    sum := int64(nums[0] + nums[1])

    for i := 2; i < len(nums); i++ {
        if int64(nums[i]) < sum {
            res = sum + int64(nums[i])
        }
        sum += int64(nums[i])
    }

    if res == 0 {
        res = -1
    }

    return res
}