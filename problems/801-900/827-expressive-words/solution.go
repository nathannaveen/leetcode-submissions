func expressiveWords(s string, words []string) int {
    arr := seperate(s)
    res := 0

    for _, word := range words {
        letters := seperate(word)

        if len(arr) != len(letters) {
            continue
        }

        i := 0
        canDo := true

        for _, l := range letters {
            if (len(arr[i]) >= 3 && l[0] == arr[i][0] && len(l) < len(arr[i])) || arr[i] == l {
                i++
                continue
            }
            canDo = false
            break
        }

        if canDo {
            res++
        }
    }

    return res
}

func seperate(s string) []string {
    arr := []string{}
    i := 0

    for j := 1; j < len(s); j++ {
        if s[j] != s[j - 1] {
            arr = append(arr, s[i : j])
            i = j
        }
    }
    arr = append(arr, s[i : len(s)])
    
    return arr
}