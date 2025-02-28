func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
    m := map[int] bool {} // whether the number in in nums
    
    for _, num := range nums {
        m[num] = true
    }

    for i := 0; i < len(moveFrom); i++ {
        delete(m, moveFrom[i])
        m[moveTo[i]] = true
    }

    res := []int{}

    for a := range m {
        res = append(res, a)
    }

    sort.Ints(res)

    return res
}