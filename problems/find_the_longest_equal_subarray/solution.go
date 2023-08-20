func longestEqualSubarray(nums []int, k int) int {
    m := map[int] int { nums[0] : 1 } // num : freq
    start := 0
    res := 1
    i := 0
    
    for i < len(nums) {
        if i - start + 1 - res > k {
            m[nums[start]]--
            start++
        } else {
            i++
            if i == len(nums) {
                break
            }
            m[nums[i]]++
            res = max(res, m[nums[i]])
        }
    }
    
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}