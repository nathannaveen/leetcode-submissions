type key struct {
    val int
    freq int
}

func topKFrequent(nums []int, k int) []int {
    arr := []key{}
    m := map[int] int {} // value : index in arr

    for _, num := range nums {
        if _, ok := m[num]; !ok {
            arr = append(arr, key{num, 0})
            m[num] = len(arr) - 1
        }
        arr[m[num]].freq++
    }

    sort.Slice(arr, func(i, j int) bool {
        return arr[i].freq > arr[j].freq
    })

    res := make([]int, k)

    for i := 0; i < k; i++ {
        res[i] = arr[i].val
    }

    return res
}