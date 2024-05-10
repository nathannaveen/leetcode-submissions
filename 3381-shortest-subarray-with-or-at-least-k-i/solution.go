func minimumSubarrayLength(nums []int, k int) int {
    min := 100

    for i := 0; i < len(nums); i++ {
        n := nums[i]
        for j := i; j < len(nums); j++ {
            n |= nums[j]

            if n >= k {
                x := j - i + 1
                if x < min {
                    min = x
                }
                break
            }
        }
    }

    if min == 100 {
        return -1
    }

    return min
}
