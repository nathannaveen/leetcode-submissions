func sumOfEncryptedInt(nums []int) int {
    res := 0

    for _, n := range nums {
        res += encrypt(n)
    }

    return res
}

func encrypt(n int) int {
    max := 0
    cnt := 0

    for n > 0 {
        x := n % 10
        if x > max {
            max = x
        }
        n /= 10
        cnt++
    }

    res := 0

    for i := 0; i < cnt; i++ {
        res = res * 10 + max
    }

    return res
}
