func isArraySpecial(nums []int) bool {
    prev := -1
    
    for _, n := range nums {
        if prev != -1 && n % 2 == prev % 2 {
            return false
        }
        prev = n
    }

    return true
}
