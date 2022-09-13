func mostFrequentEven(nums []int) int {
    max := -1
    m := map[int] int {}
    
    for _, num := range nums {
        if num % 2 == 0 {
            m[num]++
            if (m[num] > m[max]) || (m[num] == m[max] && num < max) {
                max = num
            }
        }
    }
    
    return max
}