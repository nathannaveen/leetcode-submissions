func minimumSteps(s string) int64 {
    res := int64(0)
    blacks := 0

    for _, l := range s {
        if l == '1' {
            blacks++
        } else {
            res += int64(blacks)
        }
    }

    return res
}