func maxSubarrayLength(nums []int, k int) int {
    m := map[int] int {} // map[value] all indexes that it occurs sorted
    start := 0
    cnt := 0
    max := 0

    for i := 0; i < len(nums); i++ {
        m[nums[i]]++
        cnt++

        if m[nums[i]] > k {
            for {
                m[nums[start]]--
                cnt--
                start++
                if nums[start - 1] == nums[i] {
                    break
                }
            }
        }

        if cnt > max {
            max = cnt
        }
    }

    return max
}