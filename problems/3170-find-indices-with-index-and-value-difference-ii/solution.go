type key struct {
    a, b int
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    prefix := make([]key, len(nums))
    prefixMin := make([]key, len(nums))
    max := 0
    maxIndex := -1
    min := 1000000
    minIndex := -1
    
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] >= max {
            max = nums[i]
            maxIndex = i
        }
        if nums[i] <= min {
            min = nums[i]
            minIndex = i
        }
        
        prefix[i] = key{max, maxIndex}
        prefixMin[i] = key{min, minIndex}
    }
    
    // fmt.Println(prefix)
    // fmt.Println(prefixMin)
    
    for i := 0; i < len(nums) - indexDifference; i++ {
        if abs(prefix[i + indexDifference].a - nums[i]) >= valueDifference {
            return []int{i, prefix[i + indexDifference].b}
        }
        if abs(prefixMin[i + indexDifference].a - nums[i]) >= valueDifference {
            return []int{i, prefixMin[i + indexDifference].b}
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