func minOperations(nums []int, k int) int {
    res := 0

    for _, n := range nums {
        if n < k {
            res++
        }
    }

    return res
}
