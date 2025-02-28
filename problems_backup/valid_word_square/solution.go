func validWordSquare(words []string) bool {
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            if j >= len(words) || len(words[j]) <= i || words[i][j] != words[j][i] { 
                return false 
            }
        }
    }
    return true
}