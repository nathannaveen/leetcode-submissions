func sumOfThree(num int64) []int64 {
    if (num - 3) % 3 != 0 {
        return []int64{}
    }
    
    x := (num - 3) / 3
    return []int64{ x, x + 1, x + 2 }
}