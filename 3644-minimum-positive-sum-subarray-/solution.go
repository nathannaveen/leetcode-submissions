func minimumSumSubarray(nums []int, l int, r int) int {
    prefix := []int{0}
    x := 0

    res := math.MaxInt

    for _, n := range nums {
        x += n
        prefix = append(prefix, x)
    }

    for i := 0; i < len(prefix); i++ {
        for j := i + l; j <= min(i + r, len(prefix) - 1); j++ {
            z := prefix[j] - prefix[i]
            if z > 0 && z < res {
                res = z
            }
        }
    }

    if res == math.MaxInt {
        res = -1
    }

    return res
}
