func maximumCount(nums []int) int {
    cnt := 0
    n := -1
    max := 0

    for _, num := range nums {
        if num * n > 0 {
            cnt++
        } else if num > 0 {
            if cnt > max {
                max = cnt
            }
            n = 1
            cnt = 1
        }
    }

    if cnt > max {
        max = cnt
    }

    return max
}