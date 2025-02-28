func maximizeSum(nums []int, k int) int {
    max := 0
    for _, n := range nums {
        if n > max {
            max = n
        }
    }

    return int(float64(k) / float64(2) * float64(max + max + k - 1))
}