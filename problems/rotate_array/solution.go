func rotate(nums []int, k int)  {
    start := 0
    temp := nums[0]
    visited := make(map[int] bool) // pos : bool
    for i := 0; i < len(nums); i++ {
        start = (start + k) % len(nums)
        nums[start], temp = temp, nums[start]
        // fmt.Println(start, temp, nums)
        if visited[start] {
            start += 1
            temp = nums[start]
            // fmt.Println("start", start, "temp", temp, "visited", visited)
            i--
            continue
        }
        visited[start] = true
    }
}