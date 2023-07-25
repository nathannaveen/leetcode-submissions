func peakIndexInMountainArray(arr []int) int {
    for i := 1; i < len(arr) - 1; i++ {
        if arr[i] > arr[i - 1] && arr[i] > arr[i + 1] {
            return i
        }
    }

    return 0
}