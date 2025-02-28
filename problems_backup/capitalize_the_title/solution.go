func capitalizeTitle(title string) string {
    words := strings.Split(title, " ")
    res := ""
    
    for _, word := range words {
        temp := strings.ToLower(word)
        if len(temp) > 2 {
            temp = string(temp[0] - 32) + temp[1:]
        }
        res += temp + " "
    }
    return res[:len(res) - 1]
}