func maximumEvenSplit(finalSum int64) []int64 {
    if finalSum % 2 == 1 { // odd
        return []int64{}
    }
    
    res := []int64{}
    
    i := int64(0)
    
    for finalSum > 0 {
        i += 2
        res = append(res, i)
        finalSum -= i
    }
    
    if finalSum < 0 { 
        // Remove the element from res using slice tricks
        index := (finalSum * -1) / 2
        res = append(res[: index - 1], res[index:]...)
    }
    
    return res
}