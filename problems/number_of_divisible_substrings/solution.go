func countDivisibleSubstrings(word string) int {
    values := map[string] int {
        "a" : 1, "b" : 1,
        "c" : 2, "d" : 2, "e" : 2,
        "f" : 3, "g" : 3, "h" : 3,
        "i" : 4, "j" : 4, "k" : 4,
        "l" : 5, "m" : 5, "n" : 5,
        "o" : 6, "p" : 6, "q" : 6,
        "r" : 7, "s" : 7, "t" : 7,
        "u" : 8, "v" : 8, "w" : 8,
        "x" : 9, "y" : 9, "z" : 9,
    }

    prefix := make([]int, len(word) + 1)

    res := 0

    for i, letter := range word {
        prefix[i + 1] = prefix[i] + values[string(letter)]

        for j := 0; j <= i; j++ {
            if (prefix[i + 1] - prefix[j]) % (i + 1 - j) == 0 {
                res++
            }
        }
    }

    return res
}