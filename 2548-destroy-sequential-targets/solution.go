func destroyTargets(nums []int, space int) int {
    max := 0
    res := 1000000000
    freq := map[int] int {}

    for _, n := range nums {
        freq[n % space]++
        if freq[n % space] > max {
            max = freq[n % space]
        }
    }

    for _, n := range nums {
        if freq[n % space] == max && n < res {
            res = n
        }
    }

    return res
}