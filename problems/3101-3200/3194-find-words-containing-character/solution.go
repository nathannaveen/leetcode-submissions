func findWordsContaining(words []string, x byte) []int {
    res := []int{}

    for i := 0; i < len(words); i++ {
        for _, l := range words[i] {
            if byte(l) == x {
                res = append(res, i)
                break
            }
        }
    }

    return res
}