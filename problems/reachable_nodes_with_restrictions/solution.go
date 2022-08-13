func reachableNodes(n int, edges [][]int, restricted []int) int {
    visited := map[int] bool {}
    
    for _, r := range restricted {
        visited[r] = true
    }
    
    m := map[int] []int {}
    
    res := map[int] bool {}
    
    for _, edge := range edges {
        m[edge[0]] = append(m[edge[0]], edge[1])
        m[edge[1]] = append(m[edge[1]], edge[0])
    }
    
    dfs(m, 0, visited, res)
    
    return len(res)
}

func dfs(m map[int] []int, node int, visited map[int] bool, res map[int] bool) {
    if visited[node] {
        return
    }
    
    res[node] = true
    visited[node] = true
    
    for _, n := range m[node] {
        dfs(m, n, visited, res)
    }
}