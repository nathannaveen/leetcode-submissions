func minOperations(grid [][]int, x int) int {
    res := 0
    arr := []int{}
    
    for _, row := range grid {
        arr = append(arr, row...)
    }
    
    sort.Ints(arr)
    middle := arr[len(arr) / 2]
    
    for _, n := range arr {
        if n % x != middle % x {
            return -1
        }
        res += int(math.Abs(float64((middle - n) / x))) 
    }
    return res
}