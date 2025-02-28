func wordBreak(s string, wordDict []string) []string {
    return helper(s, 0, wordDict)
}

func helper(s string, cur int, wordDict []string) []string {
    res := []string{}

    for _, word := range wordDict {
        if len(word) + cur < len(s) && s[cur : cur + len(word)] == word {
            arr := helper(s, cur + len(word), wordDict)

            for _, a := range arr {
                res = append(res, word + " " + a)
            }
        } else if len(word) + cur == len(s) && s[cur : cur + len(word)] == word {
            res = append(res, word)
        }
    }

    return res
}