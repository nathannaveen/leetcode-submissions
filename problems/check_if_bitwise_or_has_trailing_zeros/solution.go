func hasTrailingZeros(nums []int) bool {
    cnt := 0
    for _, n := range nums {
        if n % 2 == 0 {
            cnt++
        }
    }

    return cnt >= 2
}