func hasGroupsSizeX(deck []int) bool {
    m := map[int] int {}
    
    for _, d := range deck {
        m[d]++
    }
    
    min := m[deck[0]]
    
    for _, b := range m {
        min = gcd(b, min)
    }
    
    return len(m) >= 1 && min >= 2
}

func gcd(temp1 int,temp2 int) int {
    gcdNum := 0
    
    for i := 1; i <=temp1 && i <=temp2 ; i++ {
        if temp1 % i == 0 && temp2 % i == 0 {
            gcdNum=i
        }
    }
    
    return gcdNum
}  