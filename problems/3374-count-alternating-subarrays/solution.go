func countAlternatingSubarrays(nums []int) int64 {
    res := int64(0)
    prev := -1
    cnt := int64(0)

    for i := 0; i < len(nums); i++ {
        if nums[i] == prev {
            cnt = 0
        }

        cnt++
        res += cnt
        prev = nums[i]
    }

    return res
}
