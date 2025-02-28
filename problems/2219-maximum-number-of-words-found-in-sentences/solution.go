func mostWordsFound(sentences []string) int {
    res := 0
    for _, sentence := range sentences {
        words := strings.Split(sentence, " ")
        res = int(math.Max(float64(res), float64(len(words))))
    }
    return res
}