func hasIncreasingSubarrays(nums []int, k int) bool {
    incr := []bool{}

    for i := 0; i < len(nums) - k + 1; i++ {
        x := true

        for j := i + 1; j < i + k; j++ {
            if nums[j] <=nums[j - 1] {
                x = false
                break
            }
        }

        incr = append(incr, x)
        if len(incr) > k && x && incr[len(incr) - k - 1] == x {
            return true
        }
    }

    return false
}
