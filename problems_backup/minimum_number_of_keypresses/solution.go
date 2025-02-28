func minimumKeypresses(s string) int {
    arr := make([]int, 26)
    res := 0
    
    for _, letter := range s {
        arr[letter - 'a']++
    }
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })
    
    for i := 0; i < len(arr); i++ {
        res += arr[i] * int(math.Ceil(float64(i + 1) / float64(9)))
    }
    
    return res
}