func applyOperations(nums []int) []int {
    arr := []int{}
    zeros := []int{}

    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] == 0 {
            zeros = append(zeros, 0)
        } else if nums[i] == nums[i + 1] {
            arr = append(arr, nums[i] * 2)
            nums[i + 1] = 0
        } else {
            arr = append(arr, nums[i])
        }
    }

    return append(append(arr, nums[len(nums) - 1]), zeros...)
}