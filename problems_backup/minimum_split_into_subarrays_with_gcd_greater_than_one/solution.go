func minimumSplits(nums []int) int {
    n := 0
    val := 0
    
    for _, num := range nums {
        val = gcd(val, num)
        
        if val == 1 {
            n++
            val = num
        }
    }
    
    return n + 1
}

func gcd(a int, b int) int {
    for b > 0 {
        a, b = b, a % b
    }
    return a
}