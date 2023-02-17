func distinctPrimeFactors(nums []int) int {
    primes := map[int] bool {}

    for _, num := range nums {
        primeFactors(num, primes)
    }
    
    return len(primes)
}

func primeFactors(n int, primes map[int] bool) {
    cnt := 2
    for n > 1 {
        if n % cnt == 0 {
            n /= cnt
            primes[cnt] = true
        } else {
            cnt++
        }
    }
}