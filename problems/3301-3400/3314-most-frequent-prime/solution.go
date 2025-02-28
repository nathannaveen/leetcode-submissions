var primes = map[int] bool {}
var resMap = map[int] int {}
var res = 0

func mostFrequentPrime(mat [][]int) int {
    if len(primes) == 0 {
        primes = sieve(1000000)
    }
    resMap = map[int] int {}
    res = 0
    
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j++ {
            caller(mat, i, j)
        }
    }
    
    if res == 0 {
        return -1
    }
    
    return res
}

func caller(mat [][]int, i, j int) {
    move(mat, i, j, 0, 1)
    move(mat, i, j, 0, -1)
    move(mat, i, j, 1, 0)
    move(mat, i, j, -1, 0)
    move(mat, i, j, 1, 1)
    move(mat, i, j, -1, -1)
    move(mat, i, j, -1, 1)
    move(mat, i, j, 1, -1)
}

func move(mat [][]int, i, j, dirI, dirJ int) {
    n := 0
    
    I, J := i, j
    
    for I >= 0 && J >= 0 && I < len(mat) && J < len(mat[0]) {
        n = n * 10 + mat[I][J]
        if n > 10 && primes[n] {
            resMap[n]++
            if (resMap[n] > resMap[res]) || (resMap[n] == resMap[res] && n > res) {
                res = n
            }
        }
        
        I += dirI
        J += dirJ
    }
}

func sieve(n int) map[int] bool {
    integers := make([]bool, n+1)
    for i := 2; i < n+1; i++ {
            integers[i] = true
    }

    for p := 2; p*p <= n; p++ {
        if integers[p] == true {
            for i := p * 2; i <= n; i += p {
                integers[i] = false
            }
        }
    }

    primes := map[int] bool {}
    for p := 2; p <= n; p++ {
        if integers[p] == true {
            primes[p] = true
        }
    }

    return primes
}