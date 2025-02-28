func areSimilar(mat [][]int, k int) bool {
    for _, row := range mat {
        for i := 0; i < len(row); i++ {
            if row[(i + k) % len(row)] != row[i] {
                return false
            }
        }
    }

    return true
}