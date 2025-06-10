func maxDifference(s string) int {
    indexes := map[byte] int {} // char : index
    arr := []int{}
    a1, a2 := 0, 101

    for i := 0; i < len(s); i++ {
        if _, ok := indexes[s[i]]; !ok {
            arr = append(arr, 0)
            indexes[s[i]] = len(arr) - 1
        }
        arr[indexes[s[i]]]++
    }

    sort.Ints(arr)

    for i := 0; i < len(arr); i++ {
        if arr[i] % 2 == 0 {
            a2 = arr[i]
            break
        }
    }

    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] % 2 == 1 {
            a1 = arr[i]
            break
        }
    }

    return a1 - a2
}
