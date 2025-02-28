func getLucky(s string, k int) int {
    res := 0
    for _, i := range s {
        temp := int(i - 'a' + 1)
        if temp >= 10 {
            res += temp % 10
            temp /= 10
            res += temp % 10
        } else {
            res += temp
        }
    }
    
    
    for i := 1; i < k; i++ {
        temp := 0
        
        for res > 0 {
            temp += res % 10
            res /= 10
        }
        res = temp
    }
    
    return res
}

// a b c d e f g h i j k 