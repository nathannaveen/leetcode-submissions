func countDistinctIntegers(nums []int) int {
    m := map[int] bool {}
    
    for _, num := range nums {
        m[num] = true
        m[reverse(num)] = true
    }
    
    return len(m)
}

func reverse(n int) int {
    r := 0
    for n > 0 {
        r *= 10
        r += n % 10
        n /= 10
    }
    return r
}