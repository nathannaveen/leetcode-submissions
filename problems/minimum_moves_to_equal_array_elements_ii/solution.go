func minMoves2(nums []int) int {
    sort.Ints(nums)
    median := nums[len(nums) / 2]
    res := 0

    for _, num := range nums {
        res += abs(num - median)
    }

    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}