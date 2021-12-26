func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindrome(word) {
            return word
        }
    }
    return ""
}

func isPalindrome(word string) bool {
    for i := 0; i < len(word); i++ {
        if word[i] != word[len(word) - 1 - i] {
            return false
        }
    }
    return true
}