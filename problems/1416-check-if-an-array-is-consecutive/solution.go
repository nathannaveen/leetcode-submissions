func isConsecutive(nums []int) bool {
    min := 100000
    sum := 0
    visited := map[int] bool {}
    
    for _, num := range nums {
        sum += num
        if num < min {
            min = num
        }
        if visited[num] {
            return false
        }
        visited[num] = true
    }
    
    n := min + len(nums) - 1
    m := min - 1
    
    return (n * (n + 1) / 2 - m * (m + 1) / 2) == sum
}