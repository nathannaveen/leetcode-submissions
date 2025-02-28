func resultsArray(nums []int, k int) []int {
    if k == 1 {
        return nums
    }

    res := make([]int, len(nums) - k + 1)

    notConsecutive := 0

    if k != 2 {
        for i := 1; i < k - 1; i++ {
            if nums[i] != nums[i - 1] + 1 {
                notConsecutive = i + k - 1
            }
        }
    }

    for i := k - 1; i < len(nums); i++ {
        if nums[i] == nums[i - 1] + 1 && i >= notConsecutive {
            res[i - k + 1] = nums[i]
        } else if nums[i] != nums[i - 1] + 1 {
            res[i - k + 1] = -1
            notConsecutive = i + k - 1
        } else {
            res[i - k + 1] = -1
        }
    }

    return res
}
