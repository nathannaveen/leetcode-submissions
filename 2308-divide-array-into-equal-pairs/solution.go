func divideArray(nums []int) bool {
    found := map[int] bool {}
    
    for _, n := range nums {
        if found[n] {
            delete(found, n)
        } else {
            found[n] = true
        }
    }
    
    return len(found) == 0
}