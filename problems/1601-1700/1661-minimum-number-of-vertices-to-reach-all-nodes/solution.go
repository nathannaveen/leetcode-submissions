func findSmallestSetOfVertices(n int, edges [][]int) []int {
    m := make(map[int] bool)
    res := []int{}
    
    for _, edge := range edges {
        m[edge[1]] = true
    }
    for i := 0; i < n; i++ {
        if m[i] == false {
            res = append(res, i)
        }
    }
    return res
}