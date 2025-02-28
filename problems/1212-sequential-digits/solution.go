func sequentialDigits(low int, high int) []int {
    res := []int{}
    for i := len(strconv.Itoa(low)); i <= len(strconv.Itoa(high)); i++ {
        shouldBreak := false
        
        for first := 1; first <= 10 - i; first++ {
            n := 0
            for j := 0; j < i; j++ {
                n *= 10
                n += first + j
            }
            if n > high {
                shouldBreak = true
                break
            }
            if n >= low && n <= high {
                res = append(res, n)
            }
            
        }
        if shouldBreak {
            break
        }
    }
    return res
}