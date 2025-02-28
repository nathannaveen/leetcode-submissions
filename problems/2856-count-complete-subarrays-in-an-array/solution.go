func countCompleteSubarrays(nums []int) int {
    z := map[int] int {}
    
    for _, n := range nums {
        z[n]++
    }
    
    res := 0
    
    for i := 0; i < len(nums); i++ {
        res += helper(nums, i, map[int] int {}, len(z))
    }
    
    return res
}

func helper(nums []int, i int, freq map[int] int, x int) int {
    if i == len(nums) {
        return 0
    }
    
    res := 0
    
    freq[nums[i]]++

    if len(freq) == x {
        res++
    }
    
    res += helper(nums, i + 1, freq, x)
    
    return res
}