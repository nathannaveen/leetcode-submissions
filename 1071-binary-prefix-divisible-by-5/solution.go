func prefixesDivBy5(nums []int) []bool {
    n := 0
    res := make([]bool, len(nums))

    for i := 0; i < len(nums); i++ {
        n = (n * 2 + nums[i]) % 5
        res[i] = n == 0
    }

    return res
}