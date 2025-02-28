func longestAlternatingSubarray(nums []int, threshold int) int {
    if len(nums) <= 1{
        if len(nums) ==1 && nums[0] % 2 != 0 || nums[0] > threshold {
            return 0
        }
        return len(nums)
    }
    
    res := 0
    done := true
    
    for l := 0; l < len(nums); l++ {
        for r := l; r < len(nums); r++ {
            if nums[l] % 2 != 0 || nums[l] > threshold {
                continue
            }
            flag := false
            for i := l; i < r; i++ {
                if nums[i] % 2 == nums[i + 1] % 2 || nums[i] > threshold || nums[i + 1] > threshold {
                    flag = true
                    break
                }
            }
            if flag {
                continue
            }
            temp := r - l + 1
            if temp > res {
                done = false
                res = temp
            }
        }
    }
    
    if done {
        return 0
    }
    
    return res
}