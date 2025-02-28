func incremovableSubarrayCount(nums []int) int {
    res := 0

    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            increasing := true
            prev := 0

            for k := 0; k < len(nums); k++ {
                if k < i || k > j {
                    if nums[k] <= prev {
                        increasing = false
                        break
                    }
                    prev = nums[k]
                }
            }

            if increasing {
                res++
            }
        }
    }

    return res
}