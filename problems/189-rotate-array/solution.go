func rotate(nums []int, k int)  {
    start := 0
    val := nums[0]
    visited := make(map[int] bool) // pos : bool
    for i := 0; i < len(nums); i++ {
        start = (start + k) % len(nums)
        nums[start], val = val, nums[start]
        
        if visited[start] {
            start++
            val = nums[start]
            i--
            continue
        }
        visited[start] = true
    }
}