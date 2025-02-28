func maximumTripletValue(nums []int) int64 {
    max := int64(0)

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                z := int64(nums[i] - nums[j]) * int64(nums[k])

                if z > max {
                    max = z
                }
            }
        }
    }

    return max
}