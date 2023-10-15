func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            a := j - i
            b := abs(nums[j] - nums[i])
            if a >= indexDifference && b >= valueDifference {
                return []int{i, j}
            }
        }
    }
    
    return []int{-1, -1}
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}