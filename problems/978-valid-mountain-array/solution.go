func validMountainArray(arr []int) bool {
    left, right := 0, len(arr) - 1
    
    for left + 1 < len(arr) && arr[left] < arr[left + 1] {
        left++
    }
    for right > 0 && arr[right - 1] > arr[right] {
        right--
    } 
    return left > 0 && left == right && right < len(arr) - 1;
}