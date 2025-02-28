func alternateDigitSum(n int) int {
    mult := 1
    sum := 0

    for n > 0 {
        mult *= -1
        sum += mult * n % 10
        n /= 10
    }

    return mult * sum
}