func fairCandySwap(A []int, B []int) []int {
    alice, bob := 0, 0
    m := make(map[int] int)
    for _, i := range A { alice += i }
    for _, i := range B { 
        bob += i 
        m[i]++
    }
    temp := (alice - bob) / 2
    
    for _, i := range A {
        if m[i - temp] >= 1 {
            return []int{i, i - temp}
        }
    }
    
    return []int{}
}