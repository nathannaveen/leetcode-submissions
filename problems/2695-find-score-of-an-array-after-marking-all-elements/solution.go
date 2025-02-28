func findScore(nums []int) int64 {
    arr := [][]int{} // []int{ value, index }
    chosen := map[int] bool {} // index : bool

    res := int64(0)

    for i, num := range nums {
        arr = append(arr, []int{num, i})
    }

    sort.Slice(arr, func(i, j int) bool {
        if arr[i][0] == arr[j][0] {
            return arr[i][1] < arr[j][1]
        }
        return arr[i][0] < arr[j][0]
    })

    for _, a := range arr {
        if chosen[a[1]] {
            continue
        }
        res += int64(a[0])

        // don't need to mark current element because won't go over it again
        chosen[a[1] + 1] = true
        chosen[a[1] - 1] = true
    }
    
    return res
}