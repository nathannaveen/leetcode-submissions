func alternatingSubarray(nums []int) int {
    res := -1

    for i := 1; i < len(nums); i++ {
        if nums[i] - nums[i - 1] == 1 {
            res = max(res, subarray(nums, nums[i] - nums[i - 1], i) + 1)
        }
    }

    return res
}

func subarray(nums []int, add, i int) int {
    if i == len(nums) || nums[i - 1] + add != nums[i] {
        return 0
    }
    return 1 + subarray(nums, -add, i + 1)
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}