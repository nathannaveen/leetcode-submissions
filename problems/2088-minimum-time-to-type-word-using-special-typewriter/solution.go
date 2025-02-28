func minTimeToType(word string) int {
    pos := 'a'
    res := 0
    m := map[rune] int {'a' : 0, 'b' : 1, 'c' : 2, 'd' : 3, 'e' : 4, 'f' : 5, 'g' : 6, 'h' : 7, 'i' : 8, 
                        'j' : 9, 'k' : 10, 'l' : 11, 'm' : 12, 'n' : 13, 'o' : 14, 'p' : 15, 'q' : 16,
                        'r' : 17, 's' : 18, 't' : 19, 'u' : 20, 'v' : 21, 'w' : 22, 'x' : 23, 'y' : 24, 
                        'z' : 25 }
    
    for _, c := range word {
        clockwise := abs(m[c] - m[pos])
        counterClockwise := m['z'] + 1 - abs(m[pos] - m[c])
        res += min(clockwise, counterClockwise) + 1
        pos = c
    }
    
    return res
}

func min(a, b int) int {
    if a > b { return b }
    return a
}

func abs(a int) int {
    if a > 0 { return a }
    return -a
}