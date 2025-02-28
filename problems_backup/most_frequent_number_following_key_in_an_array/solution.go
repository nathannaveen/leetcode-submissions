func mostFrequent(nums []int, key int) int {
    m := map[int] int {}
    max := 0
    num := 0
    
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] == key {
            m[nums[i + 1]]++
            
            if m[nums[i + 1]] > max {
                max = m[nums[i + 1]]
                num = nums[i + 1]
            }
        }
    }
    
    return num
}