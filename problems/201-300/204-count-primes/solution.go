var sieve = make([]bool, 5000001)

func countPrimes(n int) int {
    res := 0

    for i := 2; i < n; i++ {
        if !sieve[i] {
            for j := i * 2; j < n; j += i {
                sieve[j] = true
            }
            res++
        }
    }

    return res
}