func bestHand(ranks []int, suits []byte) string {
    c := 0
    for i := 0; i < len(suits); i++ {
        if suits[i] == suits[0] {
            c++
        }
    }
    
    m := map[int] int {}
    
    max := 0
    
    for _, r := range ranks {
        m[r]++
        
        if m[r] > max {
            max = m[r]
        }
    }
    
    if c == 5 {
        return "Flush"
    }
    
    if max >= 3 {
        return "Three of a Kind"
    }
    
    if max == 2 {
        return "Pair"
    }
    
    return "High Card"
}