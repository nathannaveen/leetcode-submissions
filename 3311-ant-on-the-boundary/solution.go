func returnToBoundaryCount(nums []int) int {
    ant := 0
    res := 0

    for _, n := range nums {
        ant += n

        if ant == 0 {
            res++
        }
    }

    return res
}