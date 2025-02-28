func vowelStrings(words []string, queries [][]int) []int {
    res := []int{}

    prefixSum := []int{0}
    sum := 0
    
    vowel := map[byte] bool {
        'a' : true,
        'e' : true,
        'i' : true,
        'o' : true,
        'u' : true,
    }

    for _, word := range words {
        if vowel[word[0]] && vowel[word[len(word) - 1]] {
            sum += 1
        }
        prefixSum = append(prefixSum, sum)
    }

    for _, query := range queries {
        res = append(res, prefixSum[query[1] + 1] - prefixSum[query[0]])
    }

    return res
}

