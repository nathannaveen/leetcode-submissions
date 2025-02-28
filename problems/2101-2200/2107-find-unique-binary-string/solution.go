func findDifferentBinaryString(nums []string) string {
    stack := []string{""}
    n := len(nums[0])
    m := make(map[string] bool)
    
    for _, num := range nums {
        m[num] = true
    }
    
    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        one := pop + "1"
        zero := pop + "0"
        
        if len(one) != n {
            stack = append(stack, one, zero)
        } else if m[one] == false {
            return one
        } else if m[zero] == false {
            return zero
        }
    }
    return ""
}