func findPrimePairs(n int) [][]int {
    m := SieveOfEratosthenes(n)
    res := [][]int{}
    
    for a := range m {
        t := n - a
        
        if t >= a && m[t] {
            res = append(res, []int{a, t})
        }
    }
    
    sort.Slice(res, func(i, j int) bool {
        return res[i][0] < res[j][0]
    })
    
    return res
}

func SieveOfEratosthenes(n int) map[int]bool {
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