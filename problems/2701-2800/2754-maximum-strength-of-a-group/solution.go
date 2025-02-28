func maxStrength(nums []int) int64 {
    sort.Ints(nums)
    product := int64(1)
    prevNeg := int64(1)
    negCount := 0
    posCount := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            product *= int64(nums[i])
            posCount++
        } else if nums[i] < 0 {
            product *= int64(nums[i])
            prevNeg = int64(nums[i])
            negCount++
        }
    }
    if product <= 0 {
        product /= prevNeg
        negCount--
    }
    if posCount == 0 && negCount == 0 {
        return int64(nums[len(nums) - 1])
    }
    return product
}