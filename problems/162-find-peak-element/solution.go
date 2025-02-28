func findPeakElement(nums []int) int {
    left, right := 0, len(nums) - 1
    
    for left < right {
        mid := (left + right) / 2
        
        if nums[mid + 1] > nums[mid] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return left
}