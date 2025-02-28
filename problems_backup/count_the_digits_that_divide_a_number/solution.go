func countDigits(num int) int {
    cnt := 0
    num2 := num

    for num2 > 0 {
        if num % (num2 % 10) == 0 {
            cnt++
        }
        num2 /= 10
    }

    return cnt
}