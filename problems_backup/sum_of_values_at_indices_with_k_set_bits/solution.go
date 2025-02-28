func sumIndicesWithKSetBits(nums []int, k int) int {
    s := 0
    for i := 0; i < len(nums); i++ {
        n := 0
        x := i
        for x > 0 {
            if x & 1 == 1 {
                n++
            }
            x /= 2
        }
        if n == k {
            s += nums[i]
        }
    } 
    
    return s
}