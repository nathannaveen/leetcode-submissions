func shiftingLetters(S string, shifts []int) string {
    res := make([]rune, len(shifts))
    totalShifts := 0
    
    for i := len(shifts) - 1; i >= 0; i-- {
        totalShifts += shifts[i]
        res[i] = rune((int(S[i] - 'a') + totalShifts) % 26) + 'a'
    }

    return string(res)
}