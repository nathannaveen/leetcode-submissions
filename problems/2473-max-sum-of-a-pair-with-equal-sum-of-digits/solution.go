func maximumSum(nums []int) int {
    m := map[int] []int{} // sum of digits : values
    
    for _, n := range nums {
        sum := 0
        n2 := n
        
        for n2 > 0 {
            sum += n2 % 10
            n2 /= 10
        }
        m[sum] = append(m[sum], n)
    }
    
    max := -1
    
    for _, b := range m {
        if len(b) >= 2 {
            sort.Ints(b)
            
            sum := b[len(b) - 1] + b[len(b) - 2]
            
            if sum > max {
                max = sum
            }
        }
    }
    
    return max
}