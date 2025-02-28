func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    sums := make(map[int] int)
    res := 0
    
    for i := 0; i < len(nums1); i++ {
        for j := 0; j < len(nums2); j++ {
            sums[nums1[i] + nums2[j]]++
        }
    }
    
    for i := 0; i < len(nums3); i++ {
        for j := 0; j < len(nums4); j++ {
            if sums[ -1 * (nums3[i] + nums4[j]) ] != 0 {
                res += sums[ -1 * (nums3[i] + nums4[j]) ]
            }
        }
    }
    
    return res
}