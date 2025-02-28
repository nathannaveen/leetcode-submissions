func minOperations(nums []int) int {
    m := make(map[int]int)
    for i := range nums {
        m[nums[i]]++
    }

    res := 0
    for _, n := range m {
        for n > 1 {
            if n % 3 == 0 {
                n -= 3
            } else {
                n -= 2
            }
            res++
        }

        if n == 1 {
            return -1
        }
    }

    return res
}