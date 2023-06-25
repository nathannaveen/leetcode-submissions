func countBeautifulPairs(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            first := 0
            n := nums[i]
            for n > 0 {
                first = n % 10
                n /= 10
            }
            if GCD(first, nums[j] % 10) == 1 {
                res++
            }
        }
    }
    return res
}

func GCD(a, b int) int {
    for b != 0 {
      t := b
      b = a % b
      a = t
    }
  return a
}