func isThree(n int) bool {
    m := make(map[float64] int)
    arr := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
    for _, i := range arr {
        m[float64(i)]++
    }
    
    if m[math.Sqrt(float64(n))] == 1 {
        return true
    }
    return false
}