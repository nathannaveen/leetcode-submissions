func calculateTime(keyboard string, word string) int {
    res := 0
    cur := 0
    m := map[string] int{}
    
    for i, i2 := range keyboard {
        m[string(i2)] = i
    }
    
    for _, i := range word {
        res += int(math.Abs(float64(cur - m[string(i)])))
        cur = m[string(i)]
    }
    return res
}