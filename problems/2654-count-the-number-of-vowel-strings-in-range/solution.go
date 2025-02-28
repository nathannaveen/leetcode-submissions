func vowelStrings(words []string, left int, right int) int {
    vowel := map[byte] bool {
        'a' : true,
        'e' : true,
        'i' : true,
        'o' : true,
        'u' : true,
    }

    res := 0

    for i := left; i <= right; i++ {
        if vowel[words[i][0]] && vowel[words[i][len(words[i]) - 1]] {
            res++
        }
    }

    return res
}