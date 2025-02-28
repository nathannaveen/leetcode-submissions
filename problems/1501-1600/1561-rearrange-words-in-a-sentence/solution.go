func arrangeWords(text string) string {
    arr := strings.Split(text, " ")
    
    sort.SliceStable(arr, func(i, j int) bool {
        return len(arr[i]) < len(arr[j])
    })
    
    text = strings.ToLower(strings.Join(arr, " "))
    temp := string(int(text[0]) - 32)
    return temp + text[1:]
}