func splitNum(num int) int {
    arr := []int{}

    for num > 0 {
        arr = append(arr, num % 10)
        num /= 10
    }

    sort.Ints(arr)

    oneOrTwo := true
    one, two := 0, 0

    for _, a := range arr {
        if oneOrTwo {
            one = one * 10 + a
        } else {
            two = two * 10 + a
        }
        oneOrTwo = !oneOrTwo
    }

    return one + two
}