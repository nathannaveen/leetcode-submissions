func frequencySort(nums []int) []int {
    m := make( map[int] int ) // map[val] index in arr
    arr := [][]int{}
    res := []int{}
    
    for _, num := range nums {
        if _, ok := m[num]; !ok {
            arr = append(arr, []int{})
            m[num] = len(arr) - 1
        }
        arr[m[num]] = append(arr[m[num]], num)
    }
    
    sort.Slice(arr, func(i, j int) bool {
        if len(arr[i]) == len(arr[j]) { return arr[i][0] > arr[j][0] }
        return len(arr[i]) < len(arr[j])
    })
    
    for _, a := range arr {
        res = append(res, a...)
    }
    
    return res
}