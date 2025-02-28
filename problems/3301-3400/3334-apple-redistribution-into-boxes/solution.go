func minimumBoxes(apple []int, capacity []int) int {
    sum := 0
    sum2 := 0

    for _, a := range apple {
        sum += a
    }

    sort.Slice(capacity, func(i, j int) bool {
        return capacity[i] > capacity[j]
    })

    for i, c := range capacity {
        sum2 += c
        
        if sum2 >= sum {
            return i + 1
        }
    }

    return -1
}
