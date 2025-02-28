func semiOrderedPermutation(nums []int) int {
    cnt := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            for j := i; j >= 1; j-- {
                nums[j - 1], nums[j] = nums[j], nums[j - 1]
                cnt++
            }
            break
        }
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] == len(nums) {
            for j := i; j < len(nums) - 1; j++ {
                nums[j + 1], nums[j] = nums[j], nums[j + 1]
                cnt++
            }
            break
        }
    }
    
    return cnt
}