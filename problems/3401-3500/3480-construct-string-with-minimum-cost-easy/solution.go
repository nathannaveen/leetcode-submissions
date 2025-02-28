var dp = map[int] int {} // index in target : min cost to make the rest of the word

func minimumCost(target string, words []string, costs []int) int {
    dp = map[int] int {}
    return helper(target, 0, words, costs)
}

func helper(target string, index int, words []string, costs []int) int {
    if index == len(target) { // we have found a string that works
        return 0
    }

    if val, ok := dp[index]; ok {
        return val
    }

    val := math.MaxInt

    for i := 0; i < len(words); i++ {
        word := words[i]
        cost := costs[i]

        if index + len(word) <= len(target) && target[index : index + len(word)] == word {
            x := helper(target, index + len(word), words, costs)
            if x != -1 {
                val = min(val, x + cost)
            }
        }
    }

    if val == math.MaxInt {
        val = -1
    }

    dp[index] = val
    return val
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
