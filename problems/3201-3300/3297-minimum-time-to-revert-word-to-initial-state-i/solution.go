func minimumTimeToInitialState(word string, k int) int {
    res := int(math.Ceil(float64(len(word)) / float64(k)))

    cur := len(word) % k
    added := len(word) / k
    prefix := ""
    suffix := ""

    if cur == 0 {
        cur = k
        added -= 1
    }

    for i := 0; i < len(word); i++ {
        prefix += string(word[i])
        suffix = string(word[len(word) - 1 - i]) + suffix

        if len(suffix) == cur && added != 0 {
            if suffix == prefix {
                res = added
            }
            cur += k
            added--

            if cur >= len(word) {
                break
            }
        }
    }

    return res
}