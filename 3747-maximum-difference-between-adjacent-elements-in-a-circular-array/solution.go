func maxAdjacentDistance(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        x := abs(nums[i] - nums[(i + 1) % len(nums)])
        if x > res {
            res = x
        }
    }

    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
