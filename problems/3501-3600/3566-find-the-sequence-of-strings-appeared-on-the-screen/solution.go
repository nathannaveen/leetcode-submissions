func stringSequence(target string) []string {
    res := []string{}
    prev := ""

    for _, l := range target {
        x := 'a'
        for x != l {
            res = append(res, prev + string(x))
            x = rune(byte(int(x + 1)))
        }

        prev += string(x)

        res = append(res, prev)
    }

    return res
}
