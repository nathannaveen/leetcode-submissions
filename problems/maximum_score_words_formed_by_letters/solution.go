var dp = map[int] int {}

func maxScoreWords(words []string, letters []byte, score []int) int {
    dp = map[int] int {}

    wordLetters := map[string] map[byte] int {}
    letterFreq := map[byte] int {}
    wordToIndex := map[string] int {}

    for i, word := range words {
        wordToIndex[word] = i
        if _, ok := wordLetters[word]; !ok {
            wordLetters[word] = map[byte] int {}
        }

        for i := 0; i < len(word); i++ {
            wordLetters[word][word[i]]++
        }
    }

    for _, l := range letters {
        letterFreq[l]++
    }

    return helper(wordLetters, 0, letterFreq, score, wordToIndex)
}

func helper(wordLetters map[string] map[byte] int, used int, letterFreq map[byte] int, score []int, wordToIndex map[string] int) int {
    if val, ok := dp[used]; ok {
        return val
    }

    res := 0

    for word, letters := range wordLetters {
        if used & (1 << wordToIndex[word]) != 0 { // if word is used
            continue
        }

        n := 0
        canMake := true

        for letter, freq := range letters {
            if letterFreq[letter] < freq {
                canMake = false
                break
            }
            n += score[letter - 'a'] * freq
        }

        if canMake {
            for letter, freq := range letters {
                letterFreq[letter] -= freq
            }
            res = max(res, n + helper(wordLetters, used | (1 << wordToIndex[word]), letterFreq, score, wordToIndex))
            for letter, freq := range letters {
                letterFreq[letter] += freq
            }
        }
    }

    dp[used] = res

    return res
}