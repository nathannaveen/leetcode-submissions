var res int

func punishmentNumber(n int) int {
    res = 0
    
    for i := 1; i <= n; i++ {
        if helper(0, i * i, i) {
            // fmt.Println(i * i, i)
            res += i * i
        }
    }
    
    return res
}

func helper(sum, num, i int) bool {
    if num == 0 && sum == i {
        return true
    }
    
    res := false
    x := 0
    cnt := 1
    
    for num > 0 {
        x += num % 10 * cnt
        cnt *= 10
        num /= 10
        if helper(sum + x, num, i) {
            res = true
            break
        }
    }
    
    return res
}