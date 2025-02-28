func kthSmallest(matrix [][]int, k int) int {
    arr := make([]int, len(matrix) * len(matrix[0]))
    counter := 0
    
    for _, ints := range matrix {
        for _, i := range ints {
            arr[counter] = i
            counter++
        }
    }
    
    sort.Ints(arr)
    
    return arr[k - 1]
}