func minMaxGame(nums []int) int {
    for len(nums) > 1 {
        n := len(nums)
        c := 0
        for i := 0; i < n / 2; i++ {
            index := i * 2 - c
            val := nums[index]
            
            if (i % 2 == 0 && nums[index + 1] < val) || 
            (i % 2 == 1 && nums[index + 1] > val) {
                val = nums[index + 1]
            }
            
            nums[index] = val
            nums = append(nums[:index + 1], nums[index + 2:]...)
            c++
        }
    }
    return nums[0]
}