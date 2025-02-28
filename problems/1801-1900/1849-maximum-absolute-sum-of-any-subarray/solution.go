func maxAbsoluteSum(nums []int) int {
    sum := 0
    a, b := 0, 0
    for _, n := range nums {
        sum += n
        if sum > a {
            a = sum
        }
        if sum < b {
            b = sum
        }
    }

    return a - b
}
