func largestGoodInteger(num string) string {
    max := ""
    
    for i := 0; i < len(num) - 2; i++ {
        a := num[i]
        b := num[i + 1]
        c := num[i + 2]
        
        if a == b && b == c {
            n1, _ := strconv.Atoi(max)
            n2, _ := strconv.Atoi(num[i : i + 3])
            
            if n2 >= n1 {
                max = num[i : i + 3]
            }
        }
    }
    
    return max
}