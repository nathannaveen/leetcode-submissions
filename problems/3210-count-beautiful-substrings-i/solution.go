func beautifulSubstrings(s string, k int) int {
    res := 0

    isVowel := map[string] bool {
        "a" : true,
        "e" : true,
        "i" : true,
        "o" : true,
        "u" : true,
    }

    for i := 0; i < len(s); i++ {
        vowels, consonants := 0, 0
        for j := i; j < len(s); j++ {
            if isVowel[string(s[j])] {
                vowels++
            } else {
                consonants++
            }

            if vowels == consonants && (vowels * consonants) % k == 0 {
                res++
            }
        }
    }

    return res
}