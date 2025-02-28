func findChampion(n int, edges [][]int) int {
    arr := make([]bool, n)
    
    for _, edge := range edges {
        arr[edge[1]] = true
    }
    
    x := -1
    
    for a, b := range arr {
        if !b {
            if x != -1 {
                return -1
            }
            x = a
        }
    }
    
    return x
}