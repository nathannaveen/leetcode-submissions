func maxOperations(nums []int) int {
    res := 1
    x := nums[0] + nums[1]

    for i := 2; i < len(nums) - 1; i += 2 {
        if x != nums[i] + nums[i + 1] {
            break
        }
        res++
    }

    return res
}