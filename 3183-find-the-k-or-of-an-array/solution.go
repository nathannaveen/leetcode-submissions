func findKOr(nums []int, k int) int {
    m := map[int] int {}
    for i := 0; i < len(nums); i++ {
        cnt := 0
        x := nums[i]
        for x > 0 {
            if x & 1 == 1 {
                m[cnt]++
            }
            x >>= 1
            cnt++
        }
        if x & 1 == 1 {
            m[cnt]++
        }
    }
    
    res := 0
    
    for a, b := range m {
        if b >= k {
            res += int(math.Pow(float64(2), float64(a)))
        }
    }
    
    return res
}

/*
1 1 1
*/