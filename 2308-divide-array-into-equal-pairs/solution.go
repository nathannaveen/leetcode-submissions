func divideArray(nums []int) bool {
    numFalse := 0
    m := map[int] int {}

    for _, n := range nums {
        m[n]++
        if m[n] % 2 == 1 {
            numFalse++
        } else if m[n] > 1 {
            numFalse--
        }
    }

    return numFalse == 0
}
