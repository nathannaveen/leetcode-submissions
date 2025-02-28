func maxSum(nums []int, m int, k int) int64 {
    freq := map[int] int {}
    sum := int64(0)
    max := int64(0)

    for i := 0; i < k; i++ {
        freq[nums[i]]++
        sum += int64(nums[i])
    }

    if len(freq) >= m {
        max = sum
    }

    for i := k; i < len(nums); i++ {
        sum -= int64(nums[i - k])
        freq[nums[i - k]]--

        if freq[nums[i - k]] == 0 {
            delete(freq, nums[i - k])
        }

        sum += int64(nums[i])
        freq[nums[i]]++

        if len(freq) >= m && sum > max {
            max = sum
        }
    }

    return max
}