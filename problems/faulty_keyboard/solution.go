func finalString(s string) string {
    res := ""

    for _, l := range s {
        if l == 'i' {
            res = reverse(res)
        } else {
            res += string(l)
        }
    }

    return res
}

func reverse(s string) string {
    res := ""

    for _, l := range s {
        res = string(l) + res
    }

    return res
}