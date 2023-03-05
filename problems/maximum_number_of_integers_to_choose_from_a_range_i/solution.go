func maxCount(banned []int, n int, maxSum int) int {
    bannedMap := map[int] bool {}
    cnt := 0

    for _, ban := range banned {
        bannedMap[ban] = true
    }

    for i := 1; i <= n; i++ {
        if bannedMap[i] || maxSum - i < 0 {
            continue
        }
        maxSum -= i
        cnt++
    }

    return cnt
}