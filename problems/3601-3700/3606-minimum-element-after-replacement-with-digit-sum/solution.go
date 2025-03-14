func minElement(nums []int) int {
    min := 36
    for _, n := range nums {
        x := 0
        for n > 0 {
            x += n % 10
            n /= 10
        }
        if x < min {
            min = x
        }
    }

    return min
}
