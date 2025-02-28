func smallestValue(n int) int {
    newN := 100001
    for newN != n {
        newN = n
        n = sumPrimeFactors(n)
    }
    return newN
}

func sumPrimeFactors(n int) int {
    sum := 0
    d := 2

    for n > 1 {
        if n % d == 0 {
            n /= d
            sum += d
            continue
        }
        d++
    }

    return sum
}