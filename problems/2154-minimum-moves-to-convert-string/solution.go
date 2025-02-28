func minimumMoves(s string) int {
    res := 0 ; for i := 0; i < len(s); i++ { if s[i] == 'X' { i += 2; res++ } }; return res
}