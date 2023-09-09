func countSymmetricIntegers(low int, high int) int {
    cnt := 0
    for i := low; i <= high; i++ {
        l := len(strconv.Itoa(i))
        if l % 2 == 1 {
            continue
        }

        sum := 0
        x := i

        for j := 0; j < l / 2; j++ {
            sum += x % 10
            x /= 10
        }
        for x > 0 {
            sum -= x % 10
            x /= 10
        }

        if sum == 0 {
            cnt++
        }
    }

    return cnt
}