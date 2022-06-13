func minimumCardPickup(cards []int) int {
    m := map[int] int {}
    res := 100001
    
    for i, card := range cards {
        if _, ok := m[card]; ok && i - m[card] < res {
            res = i - m[card] + 1
        }
        m[card] = i
    }
    
    if res == 100001 {
        res = -1
    }
    
    return res
}