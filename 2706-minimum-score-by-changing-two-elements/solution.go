func minimizeSum(nums []int) int {
    sort.Ints(nums)
    one := nums[len(nums) - 1] - nums[2]
    two := nums[len(nums) - 3] - nums[0]
    three := nums[len(nums) - 2] - nums[1]

    return min(one, two, three)
}

func min(a, b, c int) int {
    if a < b {
        if a < c {
            return a
        }
        return c
    } else {
        if b < c {
            return b
        }
        return c
    }
    return -1
}