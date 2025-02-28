type WordDistance struct {
    wordsDict []string
}


func Constructor(wordsDict []string) WordDistance {
    return WordDistance{ wordsDict }
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    minimum := 30000
    one := -1
    two := -1
    
    for i, word := range this.wordsDict {
        one, minimum = findMinimum(word, word1, i, one, two, minimum)
        two, minimum = findMinimum(word, word2, i, two, one, minimum)
    }
    return minimum
}

func findMinimum(word, compareWord string, i, counter, counter2, minimum int) (int, int) {
    if word == compareWord {
        counter = i
        if counter2 != -1 && counter - counter2 < minimum {
            minimum = counter - counter2
        }
    }
    return counter, minimum
}



/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */