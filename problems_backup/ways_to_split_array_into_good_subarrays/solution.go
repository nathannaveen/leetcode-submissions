func numberOfGoodSubarraySplits(nums []int) int {
    x := 0
    res := 1
    found1 := false
    
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            if found1 {
                res *= (i - x)
                res = res % 1000000007
                x = i
            } else {
                found1 = true
                x = i
            }
        }
    }
    
    if !found1 {
        return 0
    }
    
    return res
}