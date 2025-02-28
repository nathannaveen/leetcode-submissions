func findMatrix(nums []int) [][]int {
    res := [][]int{}
    m := map[int] int {} // num : number of elements with same value

    for _, num := range nums {
        if m[num] > len(res) - 1 {
            res = append(res, []int{})
        }
        res[m[num]] = append(res[m[num]], num)
        m[num]++
    }

    return res
}