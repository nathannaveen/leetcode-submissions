func maximumSubarraySum(nums []int, k int) int64 {
    sum := int64(0)
    max := int64(0)
    duplicates := 0
    m := map[int] int {}

    for i := 0; i < k; i++ {
        sum += int64(nums[i])
        if m[nums[i]] == 1 {
            duplicates++
        }
        m[nums[i]]++
    }
    if duplicates == 0 {
        max = sum
    }

    for i := k; i < len(nums); i++ {
        m[nums[i - k]]--
        if m[nums[i - k]] == 1 {
            duplicates--
        }
        m[nums[i]]++
        if m[nums[i]] == 2 {
            duplicates++
        }
        sum += int64(nums[i]) - int64(nums[i - k])
        if duplicates == 0 && sum > max {
            max = sum
        }
    }

    return max
}