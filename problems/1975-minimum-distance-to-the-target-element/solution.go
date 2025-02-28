func getMinDistance(nums []int, target int, start int) int {
    min := 10000000
    
    for i, num := range nums {
        temp := abs(i - start)
        if num == target && temp < min {
            min = temp
        }
    }
    
    return min
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}