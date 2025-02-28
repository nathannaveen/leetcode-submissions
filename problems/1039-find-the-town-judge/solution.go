func findJudge(N int, trust [][]int) int {    
    m := make(map[int] int)
    
    for _, i := range trust {
        m[i[0]] -= len(trust)
        m[i[1]]++
    }
    
    for a, b := range m {
        if b == N - 1 { return a }
    }
    
    if N == 1 { return 1 }
    
    return -1
}