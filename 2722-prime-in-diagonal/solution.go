var primes = outputPrimes(4000000)

func diagonalPrime(nums [][]int) int {
    max := 0
    for i := 0; i < len(nums); i++ {
        if primes[nums[i][i]] && nums[i][i] > max {
            max = nums[i][i]
        }
        
        if primes[nums[i][len(nums[i]) - i - 1]] && nums[i][len(nums[i]) - i - 1] > max {
            max = nums[i][len(nums[i]) - i - 1]
        }
    }
    return max
}

func outputPrimes(limit int) map[int]bool {
    primes := map[int]bool{}
	isComposite := make([]bool, limit+1)

	for i := 2; i <= limit; i++ {
		if !isComposite[i] {
			primes[i] = true

			for j := i * i; j <= limit; j += i {
				isComposite[j] = true
			}
		}
	}

	return primes
}