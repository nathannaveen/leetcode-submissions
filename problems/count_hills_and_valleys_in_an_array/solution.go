func countHillValley(nums []int) int {
    res := 0
    
    for i := 1; i < len(nums) - 1; i++ {
        c := 0
        
        for j := i + 1; j < len(nums); j++ {
            if nums[j] > nums[i] {
                c++
                break
            } else if nums[j] < nums[i] {
                c--
                break
            }
        }
        
        if nums[i - 1] > nums[i] {
            c++
        } else if nums[i - 1] < nums[i] {
            c--
        } else {
            c = 100000
        }
        
        if c == 2 {
            res++
        } else if c == -2 {
            res++
        }
    }
    
    return res
}