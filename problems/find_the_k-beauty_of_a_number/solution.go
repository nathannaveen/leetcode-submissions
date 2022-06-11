func divisorSubstrings(num int, k int) int {
    num2 := num
    res := 0
    n := 0
    multiply := 1
    
    for i := 0; i < k - 1; i++ {
        n += (num2 % 10) * multiply
        num2 /= 10
        multiply *= 10
    }
    
    for num2 > 0 {
        n += (num2 % 10) * multiply
        
        if n != 0 && num % n == 0 {
            res++
        }
        
        n /= 10
        num2 /= 10
    }
    
    return res
}