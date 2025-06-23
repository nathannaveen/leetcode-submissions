func countPartitions(nums []int) int {
    res := 0

    totalSum := 0
    for _, n := range nums {
        totalSum += n
    }

    sum := 0

    for i, n := range nums {
        sum += n
        if i != len(nums) - 1 && (sum - totalSum - sum) % 2 == 0 {
            res++
        }
    }

    return res
}
