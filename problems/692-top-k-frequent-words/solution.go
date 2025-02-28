func topKFrequent(words []string, k int) []string {
    frequency := [][]int{}
    res := make([]string, k)

    frequency = addWordsToFrequency(words, frequency)

    sortFrequency(words, frequency)

    for i := len(frequency) - 1; i >= len(frequency)-k; i-- {
        res[len(frequency) - i - 1] = words[frequency[i][0]]
    }
    return res
}

func addWordsToFrequency(words []string, frequency [][]int) [][]int {
    /*
       adding all the words to the matrix called frequency

           * first check whether frequency has the word in it, if it does
           add one to the frequency, if it doesn't add another word to the
           array with a frequency of 1
    */
    for i, word := range words {
        contains := false

        for _, ints := range frequency {
            if words[ints[0]] == word {
                contains = true
                ints[1]++
                break
            }
        }
        if !contains {
            frequency = append(frequency, []int{i, 1})
        }
    }
    return frequency
}

func sortFrequency(words []string, frequency [][]int) {
    
    /*
    This fuction sorts the matrix called frequency.
    */
    
    for i := 0; i < len(frequency); i++ {
        if i >= 1 && (frequency[i][1] < frequency[i-1][1] ||
            (frequency[i][1] == frequency[i-1][1] && 
            words[frequency[i][0]] > words[frequency[i-1][0]])) {

            frequency[i], frequency[i-1] = frequency[i-1], frequency[i]
            i -= 2
        }
    }
}