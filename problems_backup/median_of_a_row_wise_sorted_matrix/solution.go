func matrixMedian(grid [][]int) int {
    arr := []int{}
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            arr = append(arr, grid[i][j])
        }
    }
    
    sort.Ints(arr)
    
    return arr[len(arr) / 2]
}