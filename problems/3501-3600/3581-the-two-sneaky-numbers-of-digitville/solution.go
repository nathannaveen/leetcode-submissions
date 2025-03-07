func getSneakyNumbers(nums []int) []int {
    m := map[int] bool {}
    res := []int{}

    for _, n := range nums {
        if m[n] {
            res = append(res, n)
        }
        m[n] = true
    }

    return res
}
