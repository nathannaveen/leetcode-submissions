func productExceptSelf(nums []int) (output []int) {
    m := make(map[int]int)
    for i, v := range nums {
        if mValue, ok := m[v]; ok {
            output = append(output, mValue)
            continue
        }
        
        m[v] = sum(nums[:i]) * sum(nums[i+1:])
        output = append(output, m[v])
    }
    return output
}

func sum(input []int) int {
    sum := 1
    for _, v := range input {
        if v == 0 {
            return 0
        }
        sum *= v
    }
    return sum
}