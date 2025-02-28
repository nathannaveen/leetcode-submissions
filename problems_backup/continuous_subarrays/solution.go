func continuousSubarrays(nums []int) int64 {
    res := int64(0)
    last := map[int] int {} // val : index
    x := 0
    
    for i := 0; i < len(nums); i++ {
        last[nums[i]] = i
        remove := []int{}
        a, b := nums[i] - 2, nums[i] + 2
        
        for c, d := range last {
            if c < a || c > b {
                if d + 1 > x {
                    x = d + 1
                }
                remove = append(remove, c)
            }
        }
        
        for _, c := range remove {
            delete(last, c)
        }
        
        n := i - x + 1
        res += int64(n)
    }
    
    return res
}