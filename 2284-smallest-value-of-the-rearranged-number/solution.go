func smallestNumber(num int64) int64 {
    neg := 1
    if num < 0 {
        neg = -1
        num *= -1
    }
    
    arr := strings.Split(strconv.Itoa(int(num)), "")
    
    sort.Slice(arr, func(i, j int) bool {
        if neg == -1 {
            return arr[i] > arr[j]
        }
        return arr[i] < arr[j]
    })
    
    for i := 0; i < len(arr); i++ {
        if arr[i] != "0" {
            arr[0], arr[i] = arr[i], arr[0]
            break
        }
    }
    
    temp, _ := strconv.Atoi(strings.Join(arr, ""))
    
    return int64(temp * neg)
}