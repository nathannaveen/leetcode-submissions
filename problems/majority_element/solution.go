func majorityElement(nums []int) int {
    majority := nums[0]
    counter := 0
    
    for i := 0; i < len(nums); i++ {
        if counter == 0 {
            majority = nums[i]
        }
        if majority == nums[i] {
            counter++
        } else {
            counter--
        }
    }
    
    return majority
}