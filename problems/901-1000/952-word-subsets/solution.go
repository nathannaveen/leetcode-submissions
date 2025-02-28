func wordSubsets(words1 []string, words2 []string) []string {
    arr := make([]int, 26)
    res := []string{}

    for _, word := range words2 {
        for a, b := range findFreq(word) {
            if arr[a] < b {
                arr[a] = b
            }
        }
    }

    for _, word := range words1 {
        freq := findFreq(word)
        canDo := true

        for i := 0; i < 26; i++ {
            if freq[i] < arr[i] {
                canDo = false
                break
            }
        }

        if canDo {
            res = append(res, word)
        }
    }

    return res
}

func findFreq(s string) map[int] int {
    m := map[int] int {}

    for i := 0; i < len(s); i++ {
        m[int(s[i]) - 97]++
    }

    return m
}