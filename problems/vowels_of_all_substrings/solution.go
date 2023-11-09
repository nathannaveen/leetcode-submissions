func countVowels(word string) int64 {
    res := int64(0)

    for i, l := range word {
        if strings.Contains("aeiou", string(l)) {
            res += int64((i + 1) * (len(word) - i))
        }
    }

    return res
}

/*

a b c d e f g

a - 1 2 3 4 5 6 7
0 * 7

1 2 3 4 - e - 5 6 7
4 * 3

7 + 4 + 1 = 12


*/