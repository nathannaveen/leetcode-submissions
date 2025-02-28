func sumOfTheDigitsOfHarshadNumber(x int) int {
    n := x
    res := 0

    for n > 0 {
        res += n % 10
        n /= 10
    }

    if x % res == 0 {
        return res
    }

    return -1
}
