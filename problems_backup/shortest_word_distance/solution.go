func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    minimum := 30000
    one := -1
    two := -1
    
    for i, word := range wordsDict {
        one, minimum = findMinimum(word, word1, i, one, two, minimum)
        two, minimum = findMinimum(word, word2, i, two, one, minimum)
    }
    return minimum
}

func findMinimum(word, compareWord string, i, counter, counter2, minimum int) (int, int) { // (counter, minimum)
    if word == compareWord {
        counter = i
        if counter2 != -1 && counter - counter2 < minimum {
            minimum = counter - counter2
        }
    }
    return counter, minimum
}
