func smallestEqual(nums []int) int {
    for i := range nums {
        if i % 10 == nums[i] {
            return i
        }
    }
    
    return -1
}