func findTheDifference(s string, t string) byte {
    val := 0
    
    for i := 0; i < len(s); i++ {
        val += int(t[i])
        val -= int(s[i])
    }
    val += int(t[len(t) - 1])
    return byte(val)
}