func lengthOfLastWord(s string) int {
    if len(strings.Fields(s)) == 0 { return 0 } else {return len(strings.Fields(s)[len(strings.Fields(s)) - 1]) }
}