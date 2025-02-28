func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
    i1 := Min(startPos[0], homePos[0])
    i2 := Max(startPos[0], homePos[0])
    j1 := Min(startPos[1], homePos[1])
    j2 := Max(startPos[1], homePos[1])
    
    res := 0
    
    for i := i1; i <= i2; i++ {
        res += rowCosts[i]
    }
    
    for j := j1; j <= j2; j++ {
        res += colCosts[j]
    }
    
    return res - rowCosts[startPos[0]] - colCosts[startPos[1]]
}

func Min(a, b int) int {
    if a < b {return a}
    return b
}

func Max(a, b int) int {
    if a > b { return a }
    return b
}