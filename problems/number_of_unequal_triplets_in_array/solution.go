func unequalTriplets(nums []int) int {
    m := map[int] int {}

    for _, n := range nums {
        m[n]++
    }

    n := len(nums)

    res := n * (n - 1) * (n - 2) / 6

    for _, b := range m {
        if b == 1 {
            continue
        }

        threeSame := b * (b - 1) * (b - 2) / 6
        twoSame := (n - b) * b * (b - 1) / 2

        res -= threeSame + twoSame
    }

    return res
}