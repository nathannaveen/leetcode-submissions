func decodeString(s string) string {
    return helper(s, "")
}

func helper(s string, res string) string {
    num := 0
    start := 0
    startCounter := 0
    
    for i := 0; i < len(s); i++ {
        if s[i] == '[' {
            startCounter++
            if startCounter == 1{
                start = i
            }
        } else if s[i] == ']' {
            startCounter--
            if startCounter == 0 {
                for j := 0; j < num; j++ {
                    res += helper(s[start + 1 : i], "")
                }
            }
        } else if s[i] < 97 && s[i + 1] < 97 && s[i + 2] < 97 && startCounter == 0 && 
        (s[i] != '[' && s[i] != ']' && s[i + 1] != '[' && s[i + 1] != ']' && s[i + 2] != '[' && s[i + 2] != ']')  {
            num = int(s[i] - '0') * 100 + int(s[i + 1] - '0') * 10 + int(s[i + 2] - '0')
            // fmt.Println(num, "1")
            i += 2
        } else if s[i] < 97 && s[i + 1] < 97 && startCounter == 0 && 
        (s[i] != '[' && s[i] != ']' && s[i + 1] != '[' && s[i + 1] != ']') {
            num = int(s[i] - '0') * 10 + int(s[i + 1] - '0')
            // fmt.Println(num, "2")
            i += 1
        } else if s[i] < 97 && startCounter == 0 &&
        (s[i] != '[' && s[i] != ']') {
            num = int(s[i] - '0')
            // fmt.Println(num, "3")
        } else if startCounter == 0 {
            res += string(s[i])
            // fmt.Println(string(s[i]), i)
        }
    }
    
    return res
}