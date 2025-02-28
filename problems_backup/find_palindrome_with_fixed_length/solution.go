func kthPalindrome(queries []int, intLength int) []int64 {
    minimum := int64(1)
    maximum := int64(9)
    
    for i := 1; i < int(math.Ceil(float64(intLength) / float64(2))); i++ {
        minimum *= 10
        maximum *= 10
        maximum += 9
    }
    
    res := []int64{}
    
    for _, q := range queries {
        n := minimum + int64(q) - int64(1)
        
        if n > maximum {
            res = append(res, -1)
        } else {
            res = append(res, reverseAndAdd(n, intLength))
        }
    }
    
    return res
}

func reverseAndAdd(n int64, intLength int) int64 {
    newN := n
    if intLength % 2 == 1 {
        newN /= 10
    }
    
    for newN > 0 {
        n *= 10
        n += newN % 10
        newN /= 10
    }
    
    return n
}