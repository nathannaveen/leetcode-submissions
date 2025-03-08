func reportSpam(message []string, bannedWords []string) bool {
    m := map[string] bool {}
    cnt := 0

    for _, w := range bannedWords {
        m[w] = true
    }

    for _, s := range message {
        if m[s] {
            cnt++
        }

        if cnt == 2 {
            return true
        }
    }

    return false
}
