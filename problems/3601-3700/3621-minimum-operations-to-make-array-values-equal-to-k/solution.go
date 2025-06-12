func minOperations(nums []int, k int) int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })

    res := 0
    prev := nums[0]

    for _, n := range nums {
        if n != prev {
            res++
        }
        prev = n
    }

    if nums[len(nums) - 1] < k {
        return -1
    }
    if nums[len(nums) - 1] > k {
        return res + 1
    }
    return res
}
