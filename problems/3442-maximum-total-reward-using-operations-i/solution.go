func maxTotalReward(rewardValues []int) int {
    sort.Ints(rewardValues)
    m := map[int] bool {}
    res := 0
    
    for _, v := range rewardValues {
        for z := range m {
            if v > z {
                m[v + z] = true
                if v + z > res {
                    res = v + z
                }
            }
        }
        
        m[v] = true
        if v > res {
            res = v
        }
    }
    
    return res
}
