func minOperations(nums []int) int {
    addVal := 0
    res := 0

    for i := len(nums) - 2; i >= 0; i-- {
        x := nums[i] + addVal - nums[i + 1]
        addVal += -x
        nums[i] += addVal
        if x == 0 {
            continue
        }
        res++
    }

    return res
}
