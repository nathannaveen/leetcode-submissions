func detectCapitalUse(word string) bool {
    allCap := strings.ToUpper(word)
    remainingPart := strings.ToLower(word)[1:]
    
    return  word == allCap || word[1:] == remainingPart
}