func transformArray(arr []int) []int {
    change := true
    for change {
        change = false

        arr2 := make([]int, len(arr))
        copy(arr2, arr)

        for i := 1; i < len(arr) - 1; i++ {
            left := arr[i - 1]
            right := arr[i + 1]
            
            if left < arr[i] && right < arr[i] {
                arr2[i]--
                change = true
            } else if left > arr[i] && right > arr[i] {
                arr2[i]++
                change = true
            }
        }
        arr = arr2
    }

    return arr
}