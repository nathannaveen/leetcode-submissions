func findFinalValue(nums []int, original int) int {
    m := make(map[int] bool)
    
    for _, n := range nums {
        m[n] = true
    }
    
    for m[original] {
        original <<= 1
    }
    
    return original
}