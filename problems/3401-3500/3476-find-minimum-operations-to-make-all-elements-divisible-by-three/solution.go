func minimumOperations(nums []int) int {
    res := 0

    for _, n := range nums {
        if n % 3 == 1 || n % 3 == 2 {
            res += 1
        }
    }

    return res
}
