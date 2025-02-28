func isCircularSentence(sentence string) bool {
    words := strings.Split(sentence, " ")
    words = append(words, words[0])

    letter := words[0][0]

    for _, word := range words {
        if word[0] != letter {
            return false
        }
        letter = word[len(word) - 1]
    }

    return true
}