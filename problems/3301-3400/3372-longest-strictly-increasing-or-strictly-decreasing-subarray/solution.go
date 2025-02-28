func longestMonotonicSubarray(nums []int) int {
    res := 1
    inc := 1
    dec := 1
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            dec++
            inc = 1
        } else if nums[i-1] < nums[i] {
            dec = 1
            inc++
        } else {
            inc = 1
            dec = 1
        }
        res = max(res, inc, dec)
    }
    return res
}
