func splitWordsBySeparator(words []string, separator byte) []string {
    res := []string{}

    for _, word := range words {
        s := ""
        for _, letter := range word {
            if byte(letter) == separator {
                if s != "" {
                    res = append(res, s)
                }
                s = ""
            } else {
                s += string(letter)
            }
        }
        if s != "" {
            res = append(res, s)
        }
    }

    return res
}