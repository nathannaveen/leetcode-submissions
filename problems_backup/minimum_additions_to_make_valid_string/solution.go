func addMinimum(word string) int {
    prev := byte('A')
    cnt := 0
    res := 0

    for i := 0; i < len(word); i++ {
        if word[i] <= prev {
            res += 3 - cnt
            cnt = 0
        }
        prev = word[i]
        cnt++
    }

    res += 3 - cnt

    return res
}