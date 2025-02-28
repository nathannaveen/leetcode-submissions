func maximumOddBinaryNumber(s string) string {
    arr := strings.Split(s, "")

    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })

    lastOne := 0

    for i := 0; i < len(arr); i++ {
        if arr[i] == "1" {
            lastOne = i
            continue
        }
        break
    }

    arr[lastOne], arr[len(arr) - 1] = arr[len(arr) - 1], arr[lastOne]

    return strings.Join(arr, "")
}