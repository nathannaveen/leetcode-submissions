func sumOfNumberAndReverse(num int) bool {
    for i := num / 2; i <= num; i++ {
        n := i
        i2 := 0

        for n > 0 {
            i2 = i2 * 10 + n % 10
            n /= 10
        }

        if i2 + i == num {
            return true
        }
    }

    return false
}