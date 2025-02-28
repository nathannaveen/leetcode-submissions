func edgeScore(edges []int) int {
    arr := make([][]int, len(edges)) // index, score
    
    for i := 0; i < len(edges); i++ {
        arr[i] = []int{i, 0}
    }
    
    for i, edge := range edges {
        arr[edge][1] += i
    }
    
    sort.Slice(arr, func(i, j int) bool {
        if arr[i][1] == arr[j][1] {
            return arr[i][0] < arr[j][0]
        }
        return arr[i][1] > arr[j][1]
    })
    
    return arr[0][0]
}