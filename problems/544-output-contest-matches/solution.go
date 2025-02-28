func findContestMatch(n int) string {
    arr := []string{}
    c := 0
    
    for i := 1; i <= n; i++ {
        arr = append(arr, strconv.Itoa(i))
    }
    
    for len(arr) - c != 1 {
        n := len(arr) - c
        
        for i := 0; i < n / 2; i++ {
            arr[i] = "(" + arr[i] + "," + arr[len(arr) - 1 - c] + ")"
            c++
        }
    }
    
    return arr[0]
}