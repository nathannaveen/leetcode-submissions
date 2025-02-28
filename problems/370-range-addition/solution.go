func getModifiedArray(length int, updates [][]int) []int {
    arr := make([]int, length)
    length -= 1
    
    for _, update := range updates {
        arr[update[0]] += update[2]
        if update[1] + 1 <= length {
            arr[update[1] + 1] -= update[2]
        }
    }
    
    add := 0
    for i := 0; i <= length; i++ {
        temp := add + arr[i]
        arr[i], add = temp, temp
    }
    return arr
}