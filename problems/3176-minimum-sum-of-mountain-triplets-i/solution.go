func minimumSum(nums []int) int {
    min := 100000
    
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                x := nums[i] + nums[j] + nums[k]
                if nums[j] > nums[i] && nums[j] > nums[k] && x < min {
                    min = x
                }
            }
        }
        
    }
    
    if min == 100000 {
        return -1
    }
    
    return min
}