func findTheArrayConcVal(nums []int) int64 {
    res := int64(0)

    if len(nums) % 2 == 1 {
        res += int64(nums[len(nums) / 2])
    }

    for i := 0; i < len(nums) / 2; i++ {
        n := 1

        num2 := nums[len(nums) - 1 - i]

        for num2 > 0 {
            n *= 10
            num2 /= 10
        }

        res += int64(n * nums[i] + nums[len(nums) - 1 - i])
    }

    return res
}